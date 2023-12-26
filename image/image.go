package image

import (
	"image"
	"image/color"
)

// Width and Height in pixels of a standard ZX Spectrum SCR image.
const (
	width  = 256
	height = 192
)

// Image is a ZX Spectrum compatible image implementation, which can be used
// with the standard Go image.Image interface: At(), Bounds(), ColorModel().
type Image struct {
	// Image size is fixed at 256x192 pixels
	pixels [height][width]Colour

	// Setting flash output will swap the ink/paper colours.
	enableFlashOutput bool
}

// SetFlashOutput will output an image with all FLASH colours enabled.
func (img *Image) SetFlashOutput(flash bool) {
	img.enableFlashOutput = flash
}

// Set the colour at the x/y coordinate.
func (img *Image) Set(x, y int, c Colour) {
	if x < width && y < height {
		img.pixels[y][x] = c
	}
}

// At returns the color of the pixel at the x/y coordinate.
func (img *Image) At(x, y int) color.Color {
	if x < width && y < height {
		col := img.pixels[y][x]

		// turns on the flash state if the FLASH bit was set in the SCR attribute,
		// otherwise make sure it's turned off for all colours of this image
		if img.enableFlashOutput {
			col.UseFlashColour = true
		} else {
			col.UseFlashColour = false
		}

		return col
	}
	return Colour{}
}

// Bounds returns the domain for which At can return non-zero color.
func (img *Image) Bounds() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height},
	}
}

// ColorModel returns the image's color model
// U.S. English spelling used to match the image.Image interface
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}
