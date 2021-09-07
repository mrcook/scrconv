package colour

// Colour type (a copy of color.RGBA), which can be used with the color.Color
// interface, with additional methods for converting ZX Spectrum colours.
type Colour struct {
	R, G, B, A uint8
}

func (c Colour) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

// FromAttribute converts a Spectrum pixel/attribute to RGB colour values.
// if a pixel bit is on, then the INK colour will be used, otherwise PAPER will be used.
// Colour values range from 0-7 for both INK and PAPER.
// Bit numbers:
//   7: FLASH flag
//   6: BRIGHT flag
// 3-5: PAPER colour
// 0-2: INK colour
func FromAttribute(attr uint8, pixelOn bool) Colour {
	bright := false
	if attr&0b01000000 > 0 {
		bright = true // BRIGHT is enabled
	}

	var r, g, b uint8

	if pixelOn {
		ink := attr & 0b00000111 // only the INK bits
		r, g, b = spectrumColourToRGB(ink, bright)
	} else {
		paper := (attr & 0b00111000) >> 3 // only the PAPER bits, shifted to be values 0-7
		r, g, b = spectrumColourToRGB(paper, bright)
	}

	return Colour{R: r, G: g, B: b, A: 0xFF}
}

// converts one of the ZX Spectrum's 15 colours (normal & bright) to RGB values.
func spectrumColourToRGB(zxColour uint8, bright bool) (r, g, b uint8) {
	var val uint8 = 0xEE
	if bright {
		val = 0xFF
	}

	switch zxColour {
	case 0x00:
		return 0x00, 0x00, 0x00
	case 0x01:
		return 0x00, 0x00, val
	case 0x02:
		return val, 0x00, 0x00
	case 0x03:
		return val, 0x00, val
	case 0x04:
		return 0x00, val, 0x00
	case 0x05:
		return 0x00, val, val
	case 0x06:
		return val, val, 0x00
	case 0x07:
		return val, val, val
	}
	return
}
