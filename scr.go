package scrconv

import (
	"fmt"
	"io"

	"github.com/mrcook/scrconv/colour"
	"github.com/mrcook/scrconv/image"
)

const (
	screenWidthBytes = 32 // character tiles (bytes)
	tilePosMod       = 64
)

// SCR reads a ZX Spectrum .scr file and convert to an image.
type SCR struct {
	Image *image.Image

	// offset values used for building the correct image from the raw SCR bytes.
	currentVerticalOffset   int // increments in 1/3 screen height values (64 pixels)
	currentAttributesOffset int // increments in 1/3 attribute memory size (256 bytes)
	currentPixelRowOffset   int // increment for each ZX Spectrum screen pixel row (8 pixels)

	// arrays for the raw SCR data
	memoryPixels     [6144]byte
	memoryAttributes [768]byte
}

// Convert a ZX Spectrum SCR to an image.
func (s *SCR) Convert(file io.Reader) error {
	s.Image = &image.Image{}

	if err := s.readFileBytes(file); err != nil {
		return err
	}
	s.process()
	return nil
}

// process the raw SCR bytes, generating the image data
func (s *SCR) process() {
	// process the screen in 1/3 at a time (2048 bytes) for easier conversion
	for i := 0; i < 3; i++ {
		s.resetPixelRowOffset()

		// process a sections rows in groups of 64 pixels (1/3 of total height)
		for row := 0; row < 64; row++ {
			y := s.currentVerticalOffset + row
			s.setPixelRowOffset(y)

			// set current y position for the new image
			yPos := s.currentVerticalOffset + ((y * 8) % tilePosMod) + s.currentPixelRowOffset

			for x := 0; x < screenWidthBytes; x++ {
				pixel := s.pixelsByteAt(x, y)
				attr := s.attributeAt(x, y)

				// iterate over each pixel bit, add to image, setting its
				// colour based on the corresponding attribute byte.
				for bit := 0; bit < 8; bit++ {
					pixelOn := ((pixel << bit) & 0b10000000) > 0 // is pixel enabled for current bit?

					xPos := x*8 + bit // convert byte position to correct pixel position
					s.Image.Set(xPos, yPos, colour.FromAttribute(attr, pixelOn))
				}
			}
		}
		s.currentVerticalOffset += 64
		s.currentAttributesOffset += 256
	}
}

// get pixels byte at coord
func (s *SCR) pixelsByteAt(x, y int) uint8 {
	index := screenWidthBytes * y // offset for current y position
	index += x                    // add current x position offset

	return s.memoryPixels[index]
}

// get attribute byte at coord
func (s *SCR) attributeAt(x, y int) uint8 {
	index := screenWidthBytes * y      // offset for current y position
	index = index % 256                // work in 1/3 screens (attributes memory bytes): 768/3 = 256
	index += s.currentAttributesOffset // add the current offset: 0, 256, 512
	index += x                         // add current x position offset

	return s.memoryAttributes[index]
}

// check if the offset should be increments (0-7 values) for the image y position
func (s *SCR) setPixelRowOffset(y int) {
	if y%tilePosMod > 0 && y%8 == 0 {
		s.currentPixelRowOffset++
	}
}

// reset offset for the screen current section (each 1/3 of the screen)
func (s *SCR) resetPixelRowOffset() {
	s.currentPixelRowOffset = 0
}

func (s *SCR) readFileBytes(file io.Reader) error {
	n, err := file.Read(s.memoryPixels[:])
	if err != nil {
		return err
	} else if n != 6144 {
		return fmt.Errorf("pixel data error, read only %d bytes", n)
	}

	n, err = file.Read(s.memoryAttributes[:])
	if err != nil {
		return err
	} else if n != 768 {
		return fmt.Errorf("attribute data error, read only %d bytes", n)
	}

	return nil
}
