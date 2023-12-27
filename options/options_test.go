package options_test

import (
	"fmt"
	"testing"

	"github.com/mrcook/scrconv/options"
)

func TestOptions_Validate(t *testing.T) {
	opts := options.Options{
		InFilename:   "/path/to/something.scr",
		ImageFormat:  "png",
		Scale:        2,
		WithBorder:   true,
		BorderColour: 3,
	}
	if err := opts.Validate(); err != nil {
		t.Fatalf("unexpected validation error: %s", err)
	}

	t.Run("image format validation", func(t *testing.T) {
		defer func() {
			opts.ImageFormat = "png" // reset after use
		}()

		tests := map[string]bool{
			"png": true, "jpg": true, "gif": true,
			"bmp": false, "jpeg": false, "pdf": false,
		}
		for format, valid := range tests {
			t.Run(fmt.Sprintf("format %s", format), func(t *testing.T) {
				opts.ImageFormat = format
				err := opts.Validate()
				if valid && err != nil {
					t.Errorf("unexpected validation error, got %s", err)
				} else if !valid && err == nil {
					t.Errorf("expected a validation error")
				}
			})
		}
	})

	t.Run("scale factor validation", func(t *testing.T) {
		defer func() {
			opts.Scale = 2 // reset after use
		}()

		if err := opts.Validate(); err != nil {
			t.Errorf("unexpected error, got %s", err)
		}
		opts.Scale = 0
		if err := opts.Validate(); err == nil {
			t.Errorf("expect and error")
		}
		opts.Scale = 5
		if err := opts.Validate(); err == nil {
			t.Errorf("expect and error")
		}
	})

	t.Run("border colour validation", func(t *testing.T) {
		defer func() {
			opts.BorderColour = 3 // reset after use
		}()

		opts.BorderColour = -1
		if err := opts.Validate(); err == nil {
			t.Errorf("expect and error")
		}
		opts.BorderColour = 0
		if err := opts.Validate(); err != nil {
			t.Errorf("unexpected error, got %s", err)
		}
		opts.BorderColour = 15
		if err := opts.Validate(); err != nil {
			t.Errorf("unexpected error, got %s", err)
		}
		opts.BorderColour = 16
		if err := opts.Validate(); err == nil {
			t.Errorf("expect and error")
		}
	})
}

func TestOptions_OutputFilename(t *testing.T) {
	opts := options.Options{
		InFilename:  "/path/to/something.scr",
		ImageFormat: "png",
	}

	t.Run("when no format is given", func(t *testing.T) {
		filename := opts.OutputFilename()
		if filename != "/path/to/something.png" {
			t.Errorf("unexpected filename, got '%s'", filename)
		}
	})

	t.Run("when no format is given", func(t *testing.T) {
		opts.ImageFormat = ""
		filename := opts.OutputFilename()

		if filename != "/path/to/something.new.scr" {
			t.Errorf("unexpected filename, got '%s'", filename)
		}
	})
}
