package image_test

import (
	"fmt"
	"testing"

	"github.com/mrcook/scrconv/image"
)

func TestColour(t *testing.T) {
	table := []struct {
		name    string
		attr    uint8
		pixel   bool
		r, g, b uint32
	}{
		{"black-nfg", 0b00000000, true, 0x0000, 0x0000, 0x0000},
		{"black-nbg", 0b00000000, false, 0x0000, 0x0000, 0x0000},
		{"blue-nfg", 0b00000001, true, 0x0000, 0x0000, 0xEEEE},
		{"blue-bfg", 0b01000001, true, 0x0000, 0x0000, 0xFFFF},
		{"blue-nfg", 0b00001000, false, 0x0000, 0x0000, 0xEEEE},
		{"blue-bbg", 0b01001000, false, 0x0000, 0x0000, 0xFFFF},
		{"red-nfg", 0b00000010, true, 0xEEEE, 0x0000, 0x0000},
		{"red-bfg", 0b01000010, true, 0xFFFF, 0x0000, 0x0000},
		{"red-nbg", 0b00010000, false, 0xEEEE, 0x0000, 0x0000},
		{"magenta-nfg", 0b00000011, true, 0xEEEE, 0x0000, 0xEEEE},
		{"magenta-bfg", 0b01000011, true, 0xFFFF, 0x0000, 0xFFFF},
		{"magenta-nbg", 0b00011000, false, 0xEEEE, 0x0000, 0xEEEE},
		{"green-nfg", 0b00000100, true, 0x0000, 0xEEEE, 0x0000},
		{"green-bfg", 0b01000100, true, 0x0000, 0xFFFF, 0x0000},
		{"green-nbg", 0b00100000, false, 0x0000, 0xEEEE, 0x0000},
		{"cyan-nfg", 0b00000101, true, 0x0000, 0xEEEE, 0xEEEE},
		{"cyan-bfg", 0b01000101, true, 0x0000, 0xFFFF, 0xFFFF},
		{"cyan-nbg", 0b00101000, false, 0x0000, 0xEEEE, 0xEEEE},
		{"yellow-nfg", 0b00000110, true, 0xEEEE, 0xEEEE, 0x0000},
		{"yellow-bfg", 0b01000110, true, 0xFFFF, 0xFFFF, 0x0000},
		{"yellow-nbg", 0b00110000, false, 0xEEEE, 0xEEEE, 0x0000},
		{"white-nfg", 0b00000111, true, 0xEEEE, 0xEEEE, 0xEEEE},
		{"white-bfg", 0b01000111, true, 0xFFFF, 0xFFFF, 0xFFFF},
		{"white-nbg", 0b00111000, false, 0xEEEE, 0xEEEE, 0xEEEE},
	}

	for i, col := range table {
		bright := "normal"
		if col.attr&0b01000000 > 0 {
			bright = "bright"
		}
		pixel := "paper"
		if col.pixel {
			pixel = "ink"
		}

		t.Run(fmt.Sprintf("%02d %s %s %s", i+1, bright, col.name, pixel), func(t *testing.T) {
			r, g, b, a := image.Colour{ATTR: col.attr, IsPixel: col.pixel}.RGBA()
			if r != col.r || g != col.g || b != col.b || a != 0xFFFF {
				t.Errorf("mismatch RGBA, got: %04X, %04X, %04X, %04X", r, g, b, a)
			}
		})
	}
}

func TestColour_FlashFlag(t *testing.T) {
	table := []struct {
		name    string
		attr    uint8
		pixel   bool
		r, g, b uint32
	}{
		{"black-white-n", 0b00000111, true, 0xEEEE, 0xEEEE, 0xEEEE},
		{"black-white-f", 0b10000111, true, 0x0000, 0x0000, 0x0000},
		{"green-cyan-n", 0b01100101, true, 0x0000, 0xFFFF, 0xFFFF},
		{"green-cyan-f", 0b11100101, true, 0x0000, 0xFFFF, 0x0000},
	}

	for i, col := range table {
		flash := "flash"
		if col.attr&0b01000000 > 0 {
			flash = "-----"
		}
		pixel := "paper"
		if col.pixel {
			pixel = "ink"
		}

		t.Run(fmt.Sprintf("%02d %s %s %s", i+1, flash, col.name, pixel), func(t *testing.T) {
			colour := image.Colour{
				ATTR:           col.attr,
				IsPixel:        col.pixel,
				UseFlashColour: true, // enable flash output
			}
			r, g, b, a := colour.RGBA()
			if r != col.r || g != col.g || b != col.b || a != 0xFFFF {
				t.Errorf("mismatch RGBA, got: %04X, %04X, %04X, %04X", r, g, b, a)
			}
		})
	}
}
