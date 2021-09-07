package colour_test

import (
	"fmt"
	"testing"

	"github.com/mrcook/scrconv/colour"
)

func TestColour_FromAttribute(t *testing.T) {
	table := []struct {
		name    string
		attr    uint8
		pixel   bool
		r, g, b uint32
	}{
		{"black", 0b00000000, true, 0x0000, 0x0000, 0x0000},
		{"black", 0b00000000, false, 0x0000, 0x0000, 0x0000},
		{"blue", 0b00000001, true, 0x0000, 0x0000, 0xEEEE},
		{"blue", 0b01000001, true, 0x0000, 0x0000, 0xFFFF},
		{"blue", 0b00001000, false, 0x0000, 0x0000, 0xEEEE},
		{"blue", 0b01001000, false, 0x0000, 0x0000, 0xFFFF},
		{"red", 0b00000010, true, 0xEEEE, 0x0000, 0x0000},
		{"red", 0b01000010, true, 0xFFFF, 0x0000, 0x0000},
		{"red", 0b00010000, false, 0xEEEE, 0x0000, 0x0000},
		{"magenta", 0b00000011, true, 0xEEEE, 0x0000, 0xEEEE},
		{"magenta", 0b01000011, true, 0xFFFF, 0x0000, 0xFFFF},
		{"magenta", 0b00011000, false, 0xEEEE, 0x0000, 0xEEEE},
		{"green", 0b00000100, true, 0x0000, 0xEEEE, 0x0000},
		{"green", 0b01000100, true, 0x0000, 0xFFFF, 0x0000},
		{"green", 0b00100000, false, 0x0000, 0xEEEE, 0x0000},
		{"cyan", 0b00000101, true, 0x0000, 0xEEEE, 0xEEEE},
		{"cyan", 0b01000101, true, 0x0000, 0xFFFF, 0xFFFF},
		{"cyan", 0b00101000, false, 0x0000, 0xEEEE, 0xEEEE},
		{"yellow", 0b00000110, true, 0xEEEE, 0xEEEE, 0x0000},
		{"yellow", 0b01000110, true, 0xFFFF, 0xFFFF, 0x0000},
		{"yellow", 0b00110000, false, 0xEEEE, 0xEEEE, 0x0000},
		{"white", 0b00000111, true, 0xEEEE, 0xEEEE, 0xEEEE},
		{"white", 0b01000111, true, 0xFFFF, 0xFFFF, 0xFFFF},
		{"white", 0b00111000, false, 0xEEEE, 0xEEEE, 0xEEEE},
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
			r, g, b, a := colour.FromAttribute(col.attr, col.pixel).RGBA()
			if r != col.r || g != col.g || b != col.b || a != 0xFFFF {
				t.Errorf("mismatch RGBA, got: %04X, %04X, %04X, %04X", r, g, b, a)
			}
		})
	}
}
