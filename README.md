# scrconv - ZX Spectrum SCR converter

A Go language library for converting ZX Spectrum SCR files to a Go's
`image.Image` interface, which can then be used to export to standard image
formats like PNG and JPG.

A command-line utility is also available for converting directly from a SCR
to PNG image.


## Usage as a library

Open a `.scr` file, instantiate the `SCR` struct then call the `Convert` function:

    file, err := os.Open(filename)

    scr := scrconv.SCR{}
	err = scr.Convert(file)

Once converted, `scr.Image` will be an object that responds to the Go language
`image.Image` interface, which can be used to export to a PNG image:

	f, err := os.Create(filename)

	err = png.Encode(f, img)


## Usage as a command-line app

Once you have compiled the program:

    $ scrconv -scr="/path/to/game.scr"

This command will convert the `game.scr` to a PNG image in the same directory
with the filename `game.scr.png`.


## Installation

    $ go get -u -v github.com/mrcook/scrconv/...

To install the app after manually cloning the repository:

    $ cd scrconv
    $ go install ./cmd/scrconv


## LICENSE

Copyright (c) 2021 Michael R. Cook. All rights reserved.

This work is licensed under the terms of the MIT license.
For a copy, see <https://opensource.org/licenses/MIT>.
