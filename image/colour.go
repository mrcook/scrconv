package image

// Colour represents a Spectrum pixel/attribute as an RGB colour value, and
// implements the Go color.Color interface.
type Colour struct {
	// ATTR is the original colour byte. Its bit numbers are:
	// -   7: FLASH flag
	// -   6: BRIGHT flag
	// - 3-5: PAPER colour
	// - 0-2: INK colour
	// INK/PAPER colour values range from 0-7
	ATTR uint8

	// IsPixel: indicates that the INK colour should be used
	IsPixel bool

	// UseFlashColour will output RGBA colour with the ink/paper swapped,
	// but only if the FLASH bit is also set
	UseFlashColour bool
}

// RGBA returns the RGBA colours, and respects the Go color.Color interface.
func (c Colour) RGBA() (r, g, b, a uint32) {
	ink, paper, flash := c.parseAttr()

	// swap the ink/paper colours if enabled
	if c.UseFlashColour && flash {
		ink, paper = paper, ink
	}

	// set correct pixel colour
	var col rgb
	if c.IsPixel {
		col = sinclairColourMap[ink]
	} else {
		col = sinclairColourMap[paper]
	}

	// now generate the RGBA value
	r = uint32(col.r)
	r |= r << 8
	g = uint32(col.g)
	g |= g << 8
	b = uint32(col.b)
	b |= b << 8
	a = uint32(0xFF) // the ZX Spectrum does not have transparency
	a |= a << 8
	return
}

// extracts the relevant colour data from the attribute byte.
func (c Colour) parseAttr() (uint8, uint8, bool) {
	flash := c.ATTR&0b10000000 != 0     // the FLASH flag
	bright := c.ATTR&0b01000000 != 0    // the BRIGHT flag
	paper := (c.ATTR & 0b00111000) >> 3 // the PAPER bits, shifted to be values 0-7
	ink := c.ATTR & 0b00000111          // the INK bits

	// when BRIGHT flag is set use the bright colours
	if bright {
		ink += 8
		paper += 8
	}

	return ink, paper, flash
}

// RGB represent a ZX Spectrum colour.
type rgb struct {
	r, g, b uint8
}

// The ZX Spectrum's 15 colours (normal & bright) mapped to RGB values.
var sinclairColourMap = map[uint8]rgb{
	// normal colours
	0x00: {0x00, 0x00, 0x00},
	0x01: {0x00, 0x00, 0xEE},
	0x02: {0xEE, 0x00, 0x00},
	0x03: {0xEE, 0x00, 0xEE},
	0x04: {0x00, 0xEE, 0x00},
	0x05: {0x00, 0xEE, 0xEE},
	0x06: {0xEE, 0xEE, 0x00},
	0x07: {0xEE, 0xEE, 0xEE},
	// bright colours
	0x08: {0x00, 0x00, 0x00},
	0x09: {0x00, 0x00, 0xFF},
	0x0A: {0xFF, 0x00, 0x00},
	0x0B: {0xFF, 0x00, 0xFF},
	0x0C: {0x00, 0xFF, 0x00},
	0x0D: {0x00, 0xFF, 0xFF},
	0x0E: {0xFF, 0xFF, 0x00},
	0x0F: {0xFF, 0xFF, 0xFF},
}
