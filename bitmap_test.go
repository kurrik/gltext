// Copyright 2012 The go-gl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package text

import (
	"bytes"
	"github.com/go-gl/gl"
	"github.com/go-gl/glfw"
	"image"
	"image/png"
	"testing"
)

func TestBitmap(t *testing.T) {
	var err error

	initGL(t, "bitmap")
	defer glfw.Terminate()

	// Load the same bitmap font at different scale factors.
	for i := range fonts {
		fonts[i], err = loadFont(
			loadImage(t, i+1),
			&FontConfig{
				LeftToRight,
				0, 127,
				loadGlyphs(i + 1),
			},
		)

		if err != nil {
			t.Fatal(err)
		}

		defer fonts[i].Release()
	}

	for glfw.WindowParam(glfw.Opened) > 0 {
		gl.Clear(gl.COLOR_BUFFER_BIT)

		mx, my := glfw.MousePos()
		printf(t, 10, 10, "0 1 2 3 4 5 6 7 8 9 A B C D E F")
		printf(t, float32(mx), float32(my), "%d x %d", mx, my)

		glfw.SwapBuffers()
	}
}

// loadImage loads the font image into an RGBA image.
// It optionally scales the image by the given amount.
func loadImage(t *testing.T, scale int) *image.RGBA {
	img, err := png.Decode(bytes.NewBuffer(fontImage[:]))
	if err != nil {
		t.Fatal(err)
	}
	return toRGBA(img, scale)
}

// loadGlyphs loads scaled glyph descriptors.
func loadGlyphs(scale int) Charset {
	cs := make(Charset, len(glyphData))
	copy(cs, glyphData[:])
	cs.Scale(scale)
	return cs
}

// glyphData holds glyph descriptors for fontImage.
var glyphData = [...]Glyph{
	{0, 0, 4, 8, 4}, {4, 0, 4, 8, 4}, {8, 0, 4, 8, 4}, {12, 0, 4, 8, 4},
	{16, 0, 4, 8, 4}, {20, 0, 4, 8, 4}, {24, 0, 4, 8, 4}, {28, 0, 4, 8, 4},
	{32, 0, 4, 8, 4}, {36, 0, 4, 8, 4}, {40, 0, 4, 8, 4}, {44, 0, 4, 8, 4},
	{48, 0, 4, 8, 4}, {52, 0, 4, 8, 4}, {56, 0, 4, 8, 4}, {60, 0, 4, 8, 4},
	{64, 0, 4, 8, 4}, {68, 0, 4, 8, 4}, {72, 0, 4, 8, 4}, {76, 0, 4, 8, 4},
	{80, 0, 4, 8, 4}, {84, 0, 4, 8, 4}, {88, 0, 4, 8, 4}, {92, 0, 4, 8, 4},
	{96, 0, 4, 8, 4}, {100, 0, 4, 8, 4}, {104, 0, 4, 8, 4}, {108, 0, 4, 8, 4},
	{112, 0, 4, 8, 4}, {116, 0, 4, 8, 4}, {120, 0, 4, 8, 4}, {124, 0, 4, 8, 4},

	{0, 8, 4, 8, 4}, {4, 8, 4, 8, 4}, {8, 8, 4, 8, 4}, {12, 8, 4, 8, 4},
	{16, 8, 4, 8, 4}, {20, 8, 4, 8, 4}, {24, 8, 4, 8, 4}, {28, 8, 4, 8, 4},
	{32, 8, 4, 8, 4}, {36, 8, 4, 8, 4}, {40, 8, 4, 8, 4}, {44, 8, 4, 8, 4},
	{48, 8, 4, 8, 4}, {52, 8, 4, 8, 4}, {56, 8, 4, 8, 4}, {60, 8, 4, 8, 4},
	{64, 8, 4, 8, 4}, {68, 8, 4, 8, 4}, {72, 8, 4, 8, 4}, {76, 8, 4, 8, 4},
	{80, 8, 4, 8, 4}, {84, 8, 4, 8, 4}, {88, 8, 4, 8, 4}, {92, 8, 4, 8, 4},
	{96, 8, 4, 8, 4}, {100, 8, 4, 8, 4}, {104, 8, 4, 8, 4}, {108, 8, 4, 8, 4},
	{112, 8, 4, 8, 4}, {116, 8, 4, 8, 4}, {120, 8, 4, 8, 4}, {124, 8, 4, 8, 4},

	{0, 16, 4, 8, 4}, {4, 16, 4, 8, 4}, {8, 16, 4, 8, 4}, {12, 16, 4, 8, 4},
	{16, 16, 4, 8, 4}, {20, 16, 4, 8, 4}, {24, 16, 4, 8, 4}, {28, 16, 4, 8, 4},
	{32, 16, 4, 8, 4}, {36, 16, 4, 8, 4}, {40, 16, 4, 8, 4}, {44, 16, 4, 8, 4},
	{48, 16, 4, 8, 4}, {52, 16, 4, 8, 4}, {56, 16, 4, 8, 4}, {60, 16, 4, 8, 4},
	{64, 16, 4, 8, 4}, {68, 16, 4, 8, 4}, {72, 16, 4, 8, 4}, {76, 16, 4, 8, 4},
	{80, 16, 4, 8, 4}, {84, 16, 4, 8, 4}, {88, 16, 4, 8, 4}, {92, 16, 4, 8, 4},
	{96, 16, 4, 8, 4}, {100, 16, 4, 8, 4}, {104, 16, 4, 8, 4}, {108, 16, 4, 8, 4},
	{112, 16, 4, 8, 4}, {116, 16, 4, 8, 4}, {120, 16, 4, 8, 4}, {124, 16, 4, 8, 4},

	{0, 24, 4, 8, 4}, {4, 24, 4, 8, 4}, {8, 24, 4, 8, 4}, {12, 24, 4, 8, 4},
	{16, 24, 4, 8, 4}, {20, 24, 4, 8, 4}, {24, 24, 4, 8, 4}, {28, 24, 4, 8, 4},
	{32, 24, 4, 8, 4}, {36, 24, 4, 8, 4}, {40, 24, 4, 8, 4}, {44, 24, 4, 8, 4},
	{48, 24, 4, 8, 4}, {52, 24, 4, 8, 4}, {56, 24, 4, 8, 4}, {60, 24, 4, 8, 4},
	{64, 24, 4, 8, 4}, {68, 24, 4, 8, 4}, {72, 24, 4, 8, 4}, {76, 24, 4, 8, 4},
	{80, 24, 4, 8, 4}, {84, 24, 4, 8, 4}, {88, 24, 4, 8, 4}, {92, 24, 4, 8, 4},
	{96, 24, 4, 8, 4}, {100, 24, 4, 8, 4}, {104, 24, 4, 8, 4}, {108, 24, 4, 8, 4},
	{112, 24, 4, 8, 4}, {116, 24, 4, 8, 4}, {120, 24, 4, 8, 4}, {124, 24, 4, 8, 4},
}

// fontImage holds raw, uncompressed PNG image data.
var fontImage = [...]byte{
	0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a, 0x00, 0x00, 0x00, 0x0d,
	0x49, 0x48, 0x44, 0x52, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x00, 0x20,
	0x08, 0x06, 0x00, 0x00, 0x00, 0xda, 0x22, 0x70, 0x25, 0x00, 0x00, 0x03,
	0x29, 0x49, 0x44, 0x41, 0x54, 0x78, 0xda, 0xed, 0x5b, 0xd9, 0x72, 0xc3,
	0x30, 0x08, 0xb4, 0x34, 0xf9, 0xff, 0x5f, 0x76, 0x5f, 0xaa, 0x8e, 0x86,
	0x02, 0xbb, 0x1c, 0x4e, 0xd2, 0xc6, 0xbc, 0x38, 0xae, 0x4e, 0x10, 0x5a,
	0xc1, 0x5a, 0x3d, 0x8e, 0xe3, 0x38, 0xce, 0xf3, 0x3c, 0x8f, 0x4d, 0xe4,
	0xbb, 0x94, 0x68, 0xf9, 0xfe, 0xbe, 0x7e, 0x5b, 0x4f, 0xab, 0xaf, 0xce,
	0x72, 0x6f, 0x3e, 0x48, 0x50, 0x7f, 0x1d, 0xed, 0xbd, 0xfe, 0x4e, 0x21,
	0x56, 0x39, 0x6a, 0xb7, 0x64, 0x1e, 0xb7, 0x7c, 0xb4, 0x3c, 0x9e, 0x35,
	0x10, 0xb3, 0x4b, 0xba, 0x91, 0xc7, 0x2b, 0x47, 0x88, 0xb1, 0x64, 0x8c,
	0x31, 0x22, 0xfa, 0x8d, 0x31, 0xc6, 0xde, 0x07, 0x6a, 0x2f, 0x11, 0x81,
	0x1d, 0xaf, 0x4b, 0x46, 0x14, 0xc2, 0xa4, 0x82, 0x6c, 0xfd, 0xa5, 0xd8,
	0xfa, 0x6d, 0x19, 0x69, 0xaf, 0xcb, 0x1a, 0xbe, 0x62, 0x34, 0xb4, 0xe0,
	0x48, 0x57, 0xa9, 0x9f, 0x9c, 0x53, 0xb4, 0xbd, 0xe6, 0x08, 0x5e, 0x1f,
	0xd6, 0x7c, 0x59, 0x3d, 0x1e, 0xac, 0x97, 0xca, 0x3a, 0x5e, 0x1b, 0xad,
	0xbe, 0x36, 0x81, 0xa5, 0xbc, 0x75, 0x66, 0x45, 0x9c, 0x30, 0x5a, 0x3f,
	0xb2, 0xe0, 0xda, 0x22, 0x7b, 0x08, 0x20, 0xfb, 0xd9, 0xdb, 0x6b, 0x4f,
	0x06, 0x01, 0x2c, 0x87, 0x8a, 0x6c, 0x46, 0xad, 0xff, 0xa7, 0x1d, 0x01,
	0x96, 0x02, 0xde, 0x02, 0x68, 0xc6, 0xd0, 0x0c, 0x1b, 0xa9, 0x9f, 0x41,
	0xac, 0xa8, 0x51, 0x3d, 0x7d, 0x33, 0xce, 0x75, 0xa5, 0xdc, 0x41, 0xe0,
	0x87, 0xcb, 0xed, 0x00, 0xb7, 0x70, 0xd1, 0x32, 0xca, 0x4d, 0xa3, 0x79,
	0xb2, 0x75, 0x96, 0x45, 0xcf, 0xb4, 0x4e, 0x1d, 0x33, 0xef, 0x1a, 0x9f,
	0x21, 0xe3, 0x1a, 0xf9, 0xdb, 0x2b, 0xaf, 0xd8, 0x80, 0x8d, 0xa5, 0x60,
	0xff, 0x11, 0xe2, 0xc6, 0x22, 0x25, 0x3c, 0xa2, 0xa7, 0x62, 0xe0, 0x2e,
	0x27, 0x41, 0xfa, 0x44, 0xf4, 0xf7, 0x1c, 0x98, 0xfd, 0x3b, 0x4b, 0x8c,
	0x65, 0x52, 0x60, 0x4f, 0x1f, 0x78, 0x04, 0xac, 0xe0, 0xc4, 0x4b, 0x4b,
	0xc6, 0xb7, 0xb0, 0xa9, 0x19, 0xaa, 0xbb, 0x07, 0x77, 0x57, 0x04, 0x49,
	0x5e, 0x7f, 0x28, 0xdb, 0x61, 0xe6, 0x61, 0x65, 0x12, 0x5d, 0xa9, 0xaa,
	0xf6, 0xbe, 0x07, 0x9d, 0xd2, 0x89, 0xa5, 0x3d, 0x5b, 0xbc, 0xa9, 0x02,
	0xa1, 0x4c, 0x99, 0x85, 0x30, 0x1e, 0x0d, 0xca, 0x94, 0x67, 0x10, 0x06,
	0x21, 0x14, 0x42, 0xc0, 0xcc, 0x11, 0x13, 0xa5, 0xa6, 0x33, 0xa8, 0x97,
	0x0e, 0x02, 0x35, 0x04, 0xf0, 0x16, 0xc4, 0xcb, 0xf7, 0xa3, 0x08, 0xa4,
	0x21, 0x92, 0x86, 0x46, 0x57, 0xa6, 0x55, 0x72, 0x47, 0xa1, 0xb1, 0x58,
	0xfb, 0x64, 0xd0, 0x42, 0xb6, 0xf5, 0x1c, 0x52, 0xeb, 0x77, 0xb2, 0x79,
	0xad, 0x07, 0xe1, 0x72, 0x02, 0x2c, 0x3b, 0xe5, 0x91, 0x3a, 0x1e, 0x73,
	0x58, 0xdd, 0x05, 0x15, 0xa6, 0xb1, 0x83, 0x8e, 0x66, 0x08, 0xa5, 0x08,
	0x62, 0x68, 0x0e, 0xe9, 0x1d, 0xb3, 0x50, 0xff, 0x4a, 0x00, 0xe2, 0x05,
	0x3d, 0xd1, 0xa3, 0xa5, 0x0b, 0x02, 0x23, 0x19, 0x09, 0x73, 0xe4, 0x30,
	0x7d, 0x64, 0xc7, 0xcf, 0x1e, 0x99, 0xad, 0x59, 0xc0, 0x2d, 0x1f, 0x9c,
	0xf7, 0x5b, 0x3b, 0x82, 0x2d, 0x8f, 0xee, 0xa0, 0x6a, 0x5a, 0xc9, 0xd4,
	0xeb, 0xd4, 0x8f, 0x41, 0x8c, 0xee, 0xa7, 0x65, 0x63, 0x6f, 0x87, 0xb3,
	0x69, 0xf3, 0xd4, 0x3e, 0x4a, 0x68, 0x41, 0x09, 0x5b, 0xae, 0x05, 0x46,
	0xfb, 0x3b, 0x3a, 0xd3, 0x51, 0x90, 0x57, 0x0d, 0x52, 0xab, 0xfa, 0xc9,
	0xf2, 0x8e, 0xf9, 0xa2, 0x78, 0xcb, 0xfb, 0xd6, 0xc0, 0xa6, 0xd9, 0x96,
	0xbd, 0x67, 0x65, 0x62, 0x55, 0xc5, 0xae, 0xc8, 0xe3, 0x9f, 0x25, 0xaf,
	0x1e, 0x9f, 0xc9, 0x30, 0x18, 0x07, 0x9d, 0x51, 0xa2, 0xa1, 0x8b, 0xb8,
	0xe8, 0x66, 0xf6, 0x2a, 0x8b, 0xc3, 0x44, 0xe6, 0xbb, 0x21, 0x3b, 0xe7,
	0x2b, 0x89, 0x2e, 0x74, 0xbf, 0xa0, 0x5b, 0xa8, 0x34, 0x30, 0x73, 0x09,
	0xa4, 0x9b, 0x5d, 0xab, 0xcc, 0xa7, 0xea, 0x30, 0xe8, 0xc8, 0xab, 0x1c,
	0x51, 0x08, 0xa2, 0xaf, 0x46, 0xbb, 0xc9, 0x92, 0x19, 0xaf, 0x4a, 0x25,
	0xde, 0x2d, 0x65, 0xc9, 0x5c, 0xc0, 0xb8, 0x7a, 0x11, 0x2b, 0x0e, 0x39,
	0x19, 0x96, 0x2e, 0x0a, 0x51, 0x5e, 0x7f, 0xc8, 0x20, 0xda, 0x0d, 0x19,
	0x06, 0xa2, 0x59, 0x03, 0xa0, 0xfe, 0xaa, 0xe5, 0x8c, 0x93, 0x64, 0x36,
	0x5b, 0x17, 0x57, 0x70, 0xcb, 0x1f, 0x45, 0x9e, 0x4e, 0xbd, 0x7e, 0x39,
	0x54, 0x87, 0x37, 0x65, 0x29, 0xd2, 0x28, 0x7b, 0x55, 0xf9, 0x70, 0xf2,
	0x5f, 0x1d, 0x94, 0x65, 0x1a, 0x5d, 0x3b, 0x45, 0x16, 0xa8, 0x83, 0xa2,
	0xcc, 0xee, 0xb4, 0x77, 0xa4, 0x36, 0x9f, 0x3d, 0x7e, 0xf5, 0x3e, 0x85,
	0x9a, 0x05, 0xc8, 0xef, 0xc9, 0xd9, 0x01, 0x59, 0x34, 0x40, 0x37, 0x66,
	0xa2, 0x4a, 0x77, 0x33, 0x81, 0xa8, 0xfd, 0x95, 0xcc, 0x1f, 0xc3, 0x8c,
	0x46, 0x33, 0x1c, 0xeb, 0x83, 0xd6, 0xb4, 0x98, 0x27, 0x96, 0x09, 0xb3,
	0x02, 0x42, 0x4f, 0x21, 0x74, 0x01, 0xa4, 0x7b, 0x87, 0x64, 0x98, 0x3e,
	0xd4, 0x9e, 0xfd, 0x3c, 0xdd, 0xcd, 0x04, 0x32, 0x57, 0xf4, 0x23, 0x59,
	0xc1, 0xcc, 0x42, 0x1a, 0x1b, 0x05, 0x6b, 0x86, 0xaa, 0x42, 0xe6, 0x95,
	0x37, 0x84, 0xac, 0xf9, 0x75, 0xcf, 0x1f, 0x11, 0x41, 0x91, 0xf9, 0x45,
	0x88, 0x3b, 0xb9, 0x01, 0x66, 0xd6, 0xb8, 0x88, 0x7e, 0x64, 0xb9, 0xea,
	0xcc, 0x38, 0x2c, 0xfd, 0xd9, 0x45, 0x5c, 0x55, 0xf2, 0x6e, 0x26, 0xed,
	0xed, 0xfc, 0x4f, 0xa9, 0x32, 0x13, 0xc8, 0x7a, 0xa0, 0xe5, 0x91, 0x88,
	0x57, 0x58, 0x65, 0x5d, 0x3c, 0x40, 0xf4, 0x1e, 0x41, 0x95, 0x98, 0xd1,
	0xe6, 0xaf, 0xdd, 0x61, 0xac, 0xa0, 0x27, 0xe2, 0x55, 0xaa, 0xcc, 0xe1,
	0xcd, 0x13, 0xdc, 0xf2, 0x23, 0x5f, 0x85, 0x84, 0x8c, 0x00, 0xf1, 0x44,
	0x9e, 0x43, 0x00, 0x00, 0x00, 0x00, 0x49, 0x45, 0x4e, 0x44, 0xae, 0x42,
	0x60, 0x82,
}