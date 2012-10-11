// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package text

import (
	"code.google.com/p/freetype-go/freetype"
	"code.google.com/p/freetype-go/freetype/truetype"
	"fmt"
	"github.com/go-gl/gl"
	"image"
	"io/ioutil"
	"math"

	"image/png"
	"os"
)

// Font represents a truetype font, prepared for rendering text
// to an OpenGL context.
type Font struct {
	textures []gl.Texture // Holds the texture id's.
	charset  *Charset     // Character set used to generate the font.
	scale    int32        // Font height.
	listbase uint         // Holds the first display list id.
}

// LoadFontFile loads the given truetype font and returns a Font object.
// The charset determines which rune range to use.
//
// Note: The supplied font should support the runes specified by the charset.
func LoadFontFile(file string, scale int32, cs *Charset) (font *Font, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}
	return LoadFontData(data, scale, cs)
}

// LoadFontData loads the given truetype font and returns a Font object.
// The charset determines which rune range to use.
//
// Note: The supplied font should support the runes specified by the charset.
func LoadFontData(fontData []byte, scale int32, cs *Charset) (font *Font, err error) {
	ttf, err := truetype.Parse(fontData)
	if err != nil {
		return
	}

	gb := truetype.NewGlyphBuf()

	font = new(Font)
	font.charset = cs
	font.scale = scale
	font.textures = make([]gl.Texture, cs.Len())

	gl.GenTextures(font.textures)
	font.listbase = gl.GenLists(cs.Len())

	for r := cs.Low; r <= cs.High; r++ {
		err = font.makeList(ttf, gb, r)
		if err != nil {
			return
		}
	}

	return
}

// Scale returns the font height.
func (f *Font) Scale() int32 { return f.scale }

// Charset returns the character set used to create this font.
func (f *Font) Charset() *Charset { return f.charset }

// Release cleans up all font resources.
// It can no longer be used for rendering after this call completes.
func (f *Font) Release() {
	gl.DeleteTextures(f.textures)
	gl.DeleteLists(f.listbase, f.charset.Len())

	f.charset = nil
	f.textures = nil
	f.listbase = 0
}

// Printf prints the given string at the specified coordinates.
func (f *Font) Printf(x, y float32, fs string, argv ...interface{}) {
	// Create display list indices from runes. The runes need to be offset
	// by -Charset.Low to create the correct index.
	indices := []rune(fmt.Sprintf(fs, argv...))

	for i, r := range indices {
		indices[i] = r - f.charset.Low
	}

	var vp [4]int32
	gl.GetIntegerv(gl.VIEWPORT, vp[:])

	gl.PushAttrib(gl.TRANSFORM_BIT)
	gl.MatrixMode(gl.PROJECTION)
	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Ortho(float64(vp[0]), float64(vp[2]), float64(vp[1]), float64(vp[3]), 0, 1)
	gl.PopAttrib()

	gl.PushAttrib(gl.LIST_BIT | gl.CURRENT_BIT | gl.ENABLE_BIT | gl.TRANSFORM_BIT)
	gl.MatrixMode(gl.MODELVIEW)
	gl.Disable(gl.LIGHTING)
	gl.Enable(gl.TEXTURE_2D)
	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)

	gl.ListBase(f.listbase)

	var mv [16]float32
	gl.GetFloatv(gl.MODELVIEW_MATRIX, mv[:])

	gl.PushMatrix()
	gl.LoadIdentity()
	gl.Translatef(x, (float32(vp[3]) - y - float32(f.scale)), 0)
	gl.MultMatrixf(mv[:])
	gl.CallLists(len(indices), gl.UNSIGNED_INT, indices)
	gl.PopMatrix()
	gl.PopAttrib()

	gl.PushAttrib(gl.TRANSFORM_BIT)
	gl.MatrixMode(gl.PROJECTION)
	gl.PopMatrix()
	gl.PopAttrib()
}

func pow2(n int) int { return 1 << (uint(math.Log2(float64(n))) + 1) }

// makeList makes a display list for the given glyph.
func (f *Font) makeList(ttf *truetype.Font, gb *truetype.GlyphBuf, r rune) (err error) {
	glyph := ttf.Index(r)

	err = gb.Load(ttf, f.scale, glyph, nil)
	if err != nil {
		return
	}

	// Glyph dimensions.
	glyphLeft := float32(gb.B.XMin)
	glyphTop := float32(gb.B.YMin)
	glyphWidth := float32(gb.B.XMax - gb.B.XMin)
	glyphHeight := float32(gb.B.YMax - gb.B.YMin)

	// Create power-of-two texture dimensions.
	texWidth := pow2(int(glyphWidth))
	texHeight := pow2(int(glyphHeight))

	// Create a temporary image to render to.
	rect := image.Rect(0, 0, texWidth, texHeight)
	img := image.NewGray(rect)

	// Use a freetype context to do the drawing.
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(ttf)
	c.SetFontSize(float64(f.scale))
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.White)

	// Draw the glyph.
	pt := freetype.Pt(0, int(glyphHeight))
	c.DrawString(string(r), pt)

	{
		fd, err := os.Create(fmt.Sprintf("testdata/%d.png", r))
		if err == nil {
			png.Encode(fd, img)
			fd.Close()
		}
	}

	// Index for our display list and texture. This is the same as the rune
	// value, minus the character set's lower bound.
	tex := r - f.charset.Low

	// Initialize glyph texture and render the image to it.
	f.textures[tex].Bind(gl.TEXTURE_2D)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, texWidth, texHeight,
		0, gl.LUMINANCE, gl.UNSIGNED_BYTE, img.Pix)

	// Build the display list which renders the texture to an
	// adequately positioned and scaled quad.
	gl.NewList(f.listbase+uint(tex), gl.COMPILE)
	f.textures[tex].Bind(gl.TEXTURE_2D)
	gl.Translatef(glyphLeft, 0, 0)
	gl.PushMatrix()
	gl.Translatef(0, glyphTop, 0)

	x := float64(glyphWidth) / float64(texWidth)
	y := float64(glyphHeight) / float64(texHeight)

	// Draw the quad.
	gl.Begin(gl.QUADS)
	gl.TexCoord2d(0, 0)
	gl.Vertex2f(0, glyphHeight)
	gl.TexCoord2d(0, y)
	gl.Vertex2f(0, 0)
	gl.TexCoord2d(x, y)
	gl.Vertex2f(glyphWidth, 0)
	gl.TexCoord2d(x, 0)
	gl.Vertex2f(glyphWidth, glyphHeight)
	gl.End()

	gl.PopMatrix()

	// Advance the current transformation to the next glyph location.
	metric := ttf.HMetric(f.scale, glyph)
	gl.Translatef(float32(metric.AdvanceWidth), 0, 0)

	gl.EndList()
	return
}