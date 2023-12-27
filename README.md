# scrconv - ZX Spectrum SCR converter

A CLI program for converting ZX Spectrum SCR files to normal image formats
such as PNG, GIF, and JPG.

## Usage

Once the program is installed (see below) a SCR file can be converted with
the following command:

    ./scrconv -scr="/path/to/game.scr"

By default `format=png`, `scale=1` and `border=true`, so a 320x240 PNG image is
created in the same directory as the SCR file, with the filename `game.png`.

    Usage of ./scrconv:
      -scr string
            Input .SCR filename
      -format string
            Image format: gif, jpg, png (default "png")
      -scale int
            Scale factor, max: 4 (default 1)
      -border
            Add a border to the image (default true)
      -border-colour int
            Border Colour, values: 0 - 15 (default: 0)
      -v	Show version number

### Scale

The scaling generates an image in one of the following resolutions:

    scale |   size   | + border
    ------+----------+----------
      1   |  256x192 |  320x240 (default)
      2   |  512x384 |  640x480
      3   |  768x576 |  960x720
      4   | 1024x768 | 1280x960

### Border Colour

When the `border` option is enabled, setting a `border-colour` will change
the colour of the border. The value should be one of the 16 ZX Spectrum
normal/bright colours; a value between 0-15. The default colour is black.

     # | Colour     |  #  | Colour
    ---+------------+-----+---------------
     0 | Black      |   8 | Black
     1 | Blue       |   9 | Bright Blue
     2 | Red        |  10 | Bright Red
     3 | Magenta    |  11 | Bright Magenta
     4 | Green      |  12 | Bright Green
     5 | Cyan       |  13 | Bright Cyan
     6 | Yellow     |  14 | Bright Yellow
     7 | White      |  15 | Bright White


## Installation

    go install github.com/mrcook/scrconv/cmd/scrconv@latest

To install the program after manually cloning the repository:

    cd scrconv
    go install ./cmd/scrconv


## LICENSE

Copyright (c) 2021-2023 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
