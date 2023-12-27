package scrconv

import (
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
	// TODO: create animated gif for the FLASH images
	return gif.Encode(w, img, nil)
}
