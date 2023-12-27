package options

import (
	"errors"
	"path/filepath"
	"strings"
)

type Options struct {
	InFilename       string
	ImageFormat      string
	Scale            int
	WithBorder       bool
	BackgroundColour int
}

func (o Options) OutputFilename() string {
	path := filepath.Dir(o.InFilename)
	ext := filepath.Ext(o.InFilename)
	name := strings.TrimSuffix(filepath.Base(o.InFilename), ext)

	if len(o.ImageFormat) > 0 {
		ext = "." + o.ImageFormat
	} else {
		name += ".new"
	}

	return filepath.Join(path, name+ext)
}

func (o Options) Validate() error {
	var validationErrors error

	if len(o.InFilename) == 0 {
		validationErrors = errors.Join(validationErrors, errors.New("SCR filename is missing"))
	}
	if err := o.validateFormat(); err != nil {
		validationErrors = errors.Join(validationErrors, err)
	}
	if err := o.validateScale(); err != nil {
		validationErrors = errors.Join(validationErrors, err)
	}
	if err := o.validateBgColour(); err != nil {
		validationErrors = errors.Join(validationErrors, err)
	}

	return validationErrors
}

func (o Options) validateFormat() error {
	switch o.ImageFormat {
	case "png", "jpg", "gif":
		return nil
	default:
		return errors.New("unsupported image format")
	}
}

func (o Options) validateBgColour() error {
	if o.BackgroundColour < 0 || o.BackgroundColour > 15 {
		return errors.New("background must be a ZX Spectrum colour value: 0 - 15")
	}
	return nil
}

func (o Options) validateScale() error {
	if o.Scale <= 0 || o.Scale > 4 {
		return errors.New("invalid scale factor, must be 1-4")
	}
	return nil
}
