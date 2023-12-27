package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mrcook/scrconv"
	"github.com/mrcook/scrconv/options"
)

var opts = options.Options{}

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.StringVar(&opts.InFilename, "scr", "", "Input .SCR filename")
	flag.StringVar(&opts.ImageFormat, "format", "auto", "Image format: auto, gif, jpg, png (auto=png or gif when FLASH is detected")
	flag.IntVar(&opts.Scale, "scale", 1, "Scale factor, max: 4, default: 1")
	flag.BoolVar(&opts.WithBorder, "border", true, "Add a border to the image")
	flag.IntVar(&opts.BackgroundColour, "border-colour", 0, "Border Colour, values: 0 - 15 (default 0)")
	v := flag.Bool("v", false, "Show version number")

	flag.Parse()

	if *v {
		fmt.Printf("%s v%s\n", os.Args[0], scrconv.Version)
		os.Exit(0)
	}

	if err := opts.Validate(); err != nil {
		fmt.Printf("ERROR invalid input\n%s", err)
		fmt.Println()
		fmt.Println()
		flag.Usage()
		os.Exit(2)
	}
}

func main() {
	reader, err := os.Open(opts.InFilename)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR opening SCR file: %w", err))
		os.Exit(1)
	}
	defer reader.Close()

	img, err := scrconv.ConvertToImage(reader, opts)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR reading SCR file: %w", err))
		os.Exit(1)
	}

	if opts.ImageFormat == "auto" {
		if img.HasFlashingPixels() {
			opts.ImageFormat = "gif"
		} else {
			opts.ImageFormat = "png"
		}
	}

	writer, err := os.Create(opts.OutputFilename())
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR creating PNG image file: %w", err))
		os.Exit(1)
	}
	defer writer.Close()

	switch opts.ImageFormat {
	case "png":
		if err := scrconv.ImageToPNG(writer, img); err != nil {
			fmt.Println(fmt.Errorf("ERROR convert SCR to PNG image: %w", err))
			os.Exit(1)
		}
	case "jpg":
		if err := scrconv.ImageToJPG(writer, img, 100); err != nil {
			fmt.Println(fmt.Errorf("ERROR convert SCR to JPG image: %w", err))
			os.Exit(1)
		}
	case "gif":
		if err := scrconv.ImageToGIF(writer, img); err != nil {
			fmt.Println(fmt.Errorf("ERROR convert SCR to GIF image: %w", err))
			os.Exit(1)
		}
	default:
		fmt.Println("invalid format selected")
		os.Exit(1)
	}

	fmt.Println("SCR image converted successfully")
}
