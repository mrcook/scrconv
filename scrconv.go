package scrconv

import (
	goImage "image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/mrcook/scrconv/image"
	"github.com/mrcook/scrconv/options"
)

// ConvertToImage reads the data from a SCR file and converts it to the image data.
func ConvertToImage(file io.Reader, opts options.Options) (*image.Image, error) {
	return image.FromSCR(file, opts)
}

func ImageToPNG(w io.Writer, img *image.Image) error {
	return png.Encode(w, img)
}

func ImageToJPG(w io.Writer, img *image.Image, quality int) error {
	return jpeg.Encode(w, img, &jpeg.Options{Quality: quality})
}

func ImageToGIF(w io.Writer, img *image.Image) error {
	if !img.HasFlashingPixels() {
		return gif.Encode(w, img, nil)
	}

	gifImages := &gif.GIF{
		Delay:     []int{64, 64}, // 0.64 of a second
		LoopCount: 0,             // infinite loop
	}

	// generate the base and FLASH enabled images
	for _, state := range []bool{false, true} {
		img.SetFlashOutput(state)
		gifImage := goImage.NewPaletted(img.Bounds(), image.SpectrumPalette())
		draw.Draw(gifImage, img.Bounds(), img, goImage.Point{}, draw.Src)
		gifImages.Image = append(gifImages.Image, gifImage)
	}

	return gif.EncodeAll(w, gifImages)
}
