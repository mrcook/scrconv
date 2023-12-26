package main

import (
	"flag"
	"fmt"
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
	reader, err := os.Open(*scrFilename)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR opening SCR file: %w", err))
		os.Exit(1)
	}
	defer reader.Close()

	img, err := scrconv.ReadSCR(reader)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR reading SCR file: %w", err))
		os.Exit(1)
	}

	writer, err := os.Create(*pngFilename)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR creating PNG image file: %w", err))
		os.Exit(1)
	}
	defer writer.Close()

	if err := scrconv.ImageToPNG(writer, img); err != nil {
		fmt.Println(fmt.Errorf("ERROR convert SCR to PNG image: %w", err))
		os.Exit(1)
	}

	fmt.Println("SCR image converted successfully")
}
