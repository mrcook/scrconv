package image

import (
	"image"
	"image/color"

	"github.com/mrcook/scrconv/colour"
)

// Width and Height in pixels of a standard ZX Spectrum SCR image.
const (
	Width  = 256
	Height = 192
)

// Image is an in-memory ZX Spectrum screen image implementation,
// which can be used with the image.Image interface.
// * image size is fixed at 256x192 pixels
// * pixel/attr has been replaced by a colour
// * image data stored as linear row/col data for easy conversion to modern formats
type Image struct {
	pixels [Height][Width]colour.Colour
}

// Set the colour at the x/y coordinate.
func (i *Image) Set(x, y int, c colour.Colour) {
	if x < Width && y < Height {
		i.pixels[y][x] = c
	}
}

// At returns the color of the pixel at the x/y coordinate.
func (i *Image) At(x, y int) color.Color {
	if x < Width && y < Height {
		return i.pixels[y][x]
	}
	return colour.Colour{}
}

// Bounds returns the domain for which At can return non-zero color.
func (i *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: Width, Y: Height},
	}
}

// ColorModel returns the image's color model
// U.S. English spelling used to match the image.Image interface
func (i *Image) ColorModel() color.Model {
	return color.RGBAModel
}
