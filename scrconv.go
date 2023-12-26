package scrconv

import (
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"

	"github.com/mrcook/scrconv/image"
	"github.com/mrcook/scrconv/scr"
)

// ReadSCR reads the data from a SCR file and converts it to the image data.
func ReadSCR(file io.Reader, scale int, withBorder bool) (*image.Image, error) {
	return scr.Convert(file, scale, withBorder)
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
