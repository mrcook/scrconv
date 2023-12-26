package image

import (
	"image"
	"image/color"
)

// Width and Height in pixels of a standard ZX Spectrum SCR image.
const (
	defaultWidth  = 256
	defaultHeight = 192

	// border each side = 1/8th of the image
	// with default image size: width = 32, height = 24
	borderRatio = 8
)

// Image is a ZX Spectrum compatible image implementation, which can be used
// with the standard Go image.Image interface: At(), Bounds(), ColorModel().
type Image struct {
	width, height             int
	widthBorder, heightBorder int
	pixels                    [][]Colour
	scale                     int

	// Setting flash output will swap the ink/paper colours.
	enableFlashOutput bool
}

// New return a new image using the scale value to scale the image (max x4)
func New(scale int, withBorder bool) Image {
	if scale <= 0 {
		scale = 1
	} else if scale > 4 {
		scale = 4 // set max scale level
	}

	img := Image{
		width:  defaultWidth * scale,
		height: defaultHeight * scale,
		scale:  scale,
	}
	if withBorder {
		img.widthBorder = img.width / borderRatio
		img.heightBorder = img.height / borderRatio
	}

	// initialize the pixels with the correct dimensions
	for row := 0; row < img.imageHeight(); row++ {
		var columns []Colour
		for col := 0; col < img.imageWidth(); col++ {
			columns = append(columns, Colour{})
		}
		img.pixels = append(img.pixels, columns)
	}

	return img
}

// SetFlashOutput will output an image with all FLASH colours enabled.
func (img *Image) SetFlashOutput(flash bool) {
	img.enableFlashOutput = flash
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
	y += img.heightBorder
	x += img.widthBorder

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

// imageWidth is the full width of the image, including the borders.
func (img *Image) imageWidth() int {
	return img.width + img.widthBorder*2
}

// imageHeight is the full height of the image, including the borders.
func (img *Image) imageHeight() int {
	return img.height + img.heightBorder*2
}
