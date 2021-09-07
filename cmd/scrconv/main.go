package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"os"
	"path"

	"github.com/mrcook/scrconv"
)

var (
	scrFilename  *string
	outputFormat *string
	pngFilename  *string
)

func init() {
	scrFilename = flag.String("scr", "", "Input .SCR filename")
	outputFormat = flag.String("format", "png", "Image format to output")
	pngFilename = flag.String("out", "", "Output filename (optional)")
	v := flag.Bool("v", false, "Display version number")

	flag.Parse()

	if *v {
		fmt.Printf("%s v%s\n", os.Args[0], scrconv.Version)
		os.Exit(0)
	}

	if len(*scrFilename) == 0 {
		fmt.Println("ERROR: 'scr' filename is required!")
		fmt.Println()
		flag.Usage()
		os.Exit(2)
	}

	if len(*pngFilename) == 0 {
		*pngFilename = *scrFilename
	}

	outputExtension := "." + *outputFormat
	if path.Ext(*pngFilename) != outputExtension {
		*pngFilename += outputExtension
	}
}

func main() {
	scr := scrconv.SCR{}

	if err := convertSCR(&scr, *scrFilename); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := savePNG(scr.Image, *pngFilename); err != nil {
		fmt.Printf("ERROR saving new image file: %q", err)
		os.Exit(1)
	}

	fmt.Println("SCR image converted successfully")
}

func convertSCR(scr *scrconv.SCR, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("ERROR opening SCR file: %w", err)
	}

	err = scr.Convert(file)
	if err != nil {
		return fmt.Errorf("ERROR converting SCR: %w", err)
	}

	return nil
}

func savePNG(img image.Image, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	return png.Encode(f, img)
}
