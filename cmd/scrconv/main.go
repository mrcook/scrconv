package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/mrcook/scrconv"
)

var (
	scrFilename    *string
	outputFormat   *string
	outputFilename string
	scale          *int
	withBorder     *bool
)

func init() {
	scrFilename = flag.String("scr", "", "Input .SCR filename")
	outputFormat = flag.String("format", "png", "Image format to output")
	scale = flag.Int("scale", 1, "Scale factor (default: 1, max: 4)")
	withBorder = flag.Bool("border", true, "Add a black border (default: true)")
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

	outputFilename = *scrFilename

	outputExtension := "." + *outputFormat
	if path.Ext(outputFilename) != outputExtension {
		outputFilename += outputExtension
	}
}

func main() {
	reader, err := os.Open(*scrFilename)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR opening SCR file: %w", err))
		os.Exit(1)
	}
	defer reader.Close()

	img, err := scrconv.ReadSCR(reader, *scale, *withBorder)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR reading SCR file: %w", err))
		os.Exit(1)
	}

	writer, err := os.Create(outputFilename)
	if err != nil {
		fmt.Println(fmt.Errorf("ERROR creating PNG image file: %w", err))
		os.Exit(1)
	}
	defer writer.Close()

	switch *outputFormat {
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
			fmt.Println(fmt.Errorf("ERROR convert SCR to JPG image: %w", err))
			os.Exit(1)
		}
	default:
		fmt.Println("invalid format selected")
		os.Exit(1)
	}

	fmt.Println("SCR image converted successfully")
}
