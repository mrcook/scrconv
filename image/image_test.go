package image_test

import (
	"image/color"
	"testing"

	"github.com/mrcook/scrconv/colour"
	"github.com/mrcook/scrconv/image"
)

func TestImage_SetAndAt(t *testing.T) {
	img := image.Image{}

	img.Set(10, 128, colour.Colour{R: 0x11, G: 0xCC, B: 0xBB, A: 0x99})

	t.Run("using At() with correct coordinates", func(t *testing.T) {
		at := img.At(10, 128)
		r, g, b, a := at.RGBA()

		if r != 0x1111 {
			t.Errorf("expected R to be 0x1111, got: 0x%02X", r)
		}
		if g != 0xCCCC {
			t.Errorf("expected G to be 0xCCCC, got: 0x%02X", g)
		}
		if b != 0xBBBB {
			t.Errorf("expected B to be 0xBBBB, got: 0x%02X", b)
		}
		if a != 0x9999 {
			t.Errorf("expected A to be 0x9999, got: 0x%02X", a)
		}
	})

	t.Run("when using At() with different coordinates", func(t *testing.T) {
		at := img.At(100, 256)
		r, g, b, a := at.RGBA()

		if r != 0 {
			t.Errorf("expected R to be 0, got: 0x%02X", r)
		}
		if g != 0 {
			t.Errorf("expected G to be 0, got: 0x%02X", g)
		}
		if b != 0 {
			t.Errorf("expected B to be 0, got: 0x%02X", b)
		}
		if a != 0 {
			t.Errorf("expected A to be 0, got: 0x%02X", a)
		}
	})
}

func TestImage_Bounds(t *testing.T) {
	img := image.Image{}

	x := img.Bounds().Max.X
	y := img.Bounds().Max.Y

	if x != 256 {
		t.Errorf("unexpected image bounds.X value, got %d", x)
	}
	if y != 192 {
		t.Errorf("unexpected image bounds.Y value, got %d", y)
	}
}

func TestImage_ColorModel(t *testing.T) {
	img := image.Image{}

	if img.ColorModel() != color.RGBAModel {
		t.Fatalf("unexcpeted colour model, got %s", img.ColorModel())
	}
}
