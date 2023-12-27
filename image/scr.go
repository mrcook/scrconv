package image

import (
	"fmt"
	"io"

	"github.com/mrcook/scrconv/options"
)

// FromSCR a ZX Spectrum SCR to an Image representation.
func FromSCR(file io.Reader, opts options.Options) (*Image, error) {
	s := scr{}

	if err := s.readFileBytes(file); err != nil {
		return nil, err
	}

	img := New(opts)

	// process the screen in 1/3 at a time (2048 bytes) for easier conversion
	for i := 0; i < 3; i++ {
		// reset offset for the screen current section (each 1/3 of the screen)
		s.pixelRowOffset = 0

		// process a section's rows in groups of 64 pixels (1/3 of total height)
		for row := 0; row < 64; row++ {
			y := s.verticalOffset + row
			s.setPixelRowOffset(y)

			// set current y position for the new image
			yPos := s.verticalOffset + ((y * 8) % tilePosMod) + s.pixelRowOffset

			for x := 0; x < screenWidthBytes; x++ {
				pixel := s.pixelsByteAt(x, y)
				attr := s.attributeAt(x, y)

				// if one attribute has the FLASH bit set
				if !img.hasFlashingPixels {
					if attr&0b10000000 != 0 {
						img.hasFlashingPixels = true
					}
				}

				// iterate over each pixel bit, add to image, setting its
				// colour based on the corresponding attribute byte.
				for bit := 0; bit < 8; bit++ {
					// is pixel enabled for current bit?
					isPixel := ((pixel << bit) & 0b10000000) > 0

					xPos := x*8 + bit // convert byte position to correct pixel position
					img.Set(xPos, yPos, Colour{ATTR: attr, IsPixel: isPixel})
				}
			}
		}
		s.verticalOffset += 64
		s.attributesOffset += 256
	}

	return &img, nil
}

const (
	screenWidthBytes = 32 // character tiles (bytes)
	tilePosMod       = 64
)

// SCR reads a ZX Spectrum .scr file and convert to an image.
type scr struct {
	// arrays for the raw SCR data
	pixels     [6144]byte
	attributes [768]byte

	// offset values used for building the correct image from the raw SCR bytes.
	verticalOffset   int // increments in 1/3 screen height values (64 pixels)
	attributesOffset int // increments in 1/3 attribute memory size (256 bytes)
	pixelRowOffset   int // increment for each ZX Spectrum screen pixel row (8 pixels)
}

// get pixel byte at coordinate
func (s *scr) pixelsByteAt(x, y int) uint8 {
	index := screenWidthBytes * y // offset for current y position
	index += x                    // add current x position offset

	return s.pixels[index]
}

// get attribute byte at coordinate
func (s *scr) attributeAt(x, y int) uint8 {
	index := screenWidthBytes * y // offset for current y position
	index = index % 256           // work in 1/3 screens (attributes memory bytes): 768/3 = 256
	index += s.attributesOffset   // add the current offset: 0, 256, 512
	index += x                    // add current x position offset

	return s.attributes[index]
}

// check if the offset should be increments (0-7 values) for the image y position
func (s *scr) setPixelRowOffset(y int) {
	if y%tilePosMod > 0 && y%8 == 0 {
		s.pixelRowOffset++
	}
}

func (s *scr) readFileBytes(file io.Reader) error {
	n, err := file.Read(s.pixels[:])
	if err != nil {
		return err
	} else if n != 6144 {
		return fmt.Errorf("pixel error, only %d bytes read", n)
	}

	n, err = file.Read(s.attributes[:])
	if err != nil {
		return err
	} else if n != 768 {
		return fmt.Errorf("attribute error, only %d bytes read", n)
	}

	return nil
}
