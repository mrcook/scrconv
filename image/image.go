package image

import (
	"image"
	"image/color"

	"github.com/mrcook/scrconv/options"
)

// Dimension of a standard ZX Spectrum SCR image in pixels.
const (
	defaultWidth        = 256
	defaultHeight       = 192
	defaultWidthBorder  = 32 // border each side = 1/8th of the image width
	defaultHeightBorder = 24 // border each side = 1/8th of the image height
)

// Image is a ZX Spectrum compatible image implementation, which can be used
// with the standard Go image.Image interface: At(), Bounds(), ColorModel().
type Image struct {
	enableFlashOutput bool       // when enabled will swap the ink/paper colours
	hasFlashingPixels bool       // set when a pixel has the FLASH bit set
	scale             int        // scale factor: 1-4
	bordered          bool       // should the image include a border
	borderColour      Colour     // if border enabled what colour? default: black
	pixels            [][]Colour // the image pixels
}

// New returns a new image with the given options.
func New(opts options.Options) Image {
	img := Image{
		scale:    opts.Scale,
		bordered: opts.WithBorder,
	}
	img.setBackgroundColour(opts.BackgroundColour)

	// initialize the pixels with the correct dimensions (with scaling and borders)
	for row := 0; row < img.imageHeight(); row++ {
		var columns []Colour
		for col := 0; col < img.imageWidth(); col++ {
			columns = append(columns, img.borderColour)
		}
		img.pixels = append(img.pixels, columns)
	}

	return img
}

// SetFlashOutput will output an image with all FLASH colours on.
func (img *Image) SetFlashOutput(flash bool) {
	img.enableFlashOutput = flash
}

// HasFlashingPixels returns true when the image has at least one pixel with the FLASH bit set.
func (img *Image) HasFlashingPixels() bool {
	return img.hasFlashingPixels
}

// Set the colour at the x/y coordinate, applying the borders and any scaling.
func (img *Image) Set(x, y int, c Colour) {
	if x >= defaultWidth || y >= defaultHeight {
		return
	}

	// apply scaling to the starting point
	y *= img.scale
	x *= img.scale

	// add padding for the left/top borders
	y += img.scaledHeightBorder()
	x += img.scaledWidthBorder()

	// generate the correct number of pixels for the scaling factor
	for row := 0; row < img.scale; row++ {
		for col := 0; col < img.scale; col++ {
			img.pixels[y+row][x+col] = c
		}
	}
}

// At returns the color of the pixel at the x/y coordinate.
func (img *Image) At(x, y int) color.Color {
	if x < img.imageWidth() && y < img.imageHeight() {
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
		Max: image.Point{X: img.imageWidth(), Y: img.imageHeight()},
	}
}

// ColorModel returns the image's color model
// U.S. English spelling used to match the image.Image interface
func (img *Image) ColorModel() color.Model {
	return color.RGBAModel
}

// imageWidth is the full width of the image, including the borders, with scaling applied.
func (img *Image) imageWidth() int {
	return defaultWidth*img.scale + img.scaledWidthBorder()*2
}

// imageHeight is the full height of the image, including the borders, with scaling applied.
func (img *Image) imageHeight() int {
	return defaultHeight*img.scale + img.scaledHeightBorder()*2
}

// scaledWidthBorder is border size, with scaling applied.
func (img *Image) scaledWidthBorder() int {
	if img.bordered {
		return defaultWidthBorder * img.scale
	}
	return 0
}

// scaledHeightBorder is border size, with scaling applied.
func (img *Image) scaledHeightBorder() int {
	if img.bordered {
		return defaultHeightBorder * img.scale
	}
	return 0
}

func (img *Image) setBackgroundColour(colour int) {
	if colour <= 0x00 || colour > 0x0F {
		return
	}

	var attr = uint8(colour)

	// set as a BRIGHT colour
	if attr > 0x07 {
		attr -= 8
		attr |= 1 << 6
	}

	img.borderColour = Colour{ATTR: attr, IsPixel: true}
}
