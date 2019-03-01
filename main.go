// fov: Calculate horizontal or vertical FOV values for a given aspect ratio
//
// Copyright © 2018 Hugo Locurcio and contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli"
)

type fov struct {
	horizontal  float64 // The "old" horizontal field of view
	vertical    float64 // The "old" vertical field of view
	aspectRatio float64 // The "old" aspect ratio

	// The following fields are only used if 3 arguments were given on the command line
	newHorizontal  float64 // The "new" horizontal field of view
	newVertical    float64 // The "new" vertical field of view
	newAspectRatio float64 // The "new" aspect ratio to scale the FOV to
}

// fractionToFloat converts a fraction in 'x:y' or 'x/y' notation to a floating-point value.
// It will return an error and exit if the input is not a fraction.
func fractionToFloat(fraction string) float64 {
	var terms []string

	if strings.Count(fraction, ":") == 1 {
		terms = strings.Split(fraction, ":")
	} else if strings.Count(fraction, "/") == 1 {
		terms = strings.Split(fraction, "/")
	} else {
		fmt.Fprintln(
			color.Output,
			color.HiRedString("Error:"),
			"Bad format for aspect ratio; the ratio must have a 'x:y' or 'x/y' format.",
		)
		os.Exit(1)
	}

	numerator, _ := strconv.ParseFloat(terms[0], 64)
	denominator, _ := strconv.ParseFloat(terms[1], 64)

	return numerator / denominator
}

// degreeString returns a string with the degree symbol appended, e.g. "90°".
func degreeString(float float64) string {
	return strconv.FormatFloat(float, 'f', 2, 64) + "°"
}

func main() {
	app := cli.NewApp()
	app.Name = "fov"
	app.Version = "0.2.1"
	app.Usage = "Calculate horizontal or vertical FOV values for a given aspect ratio"
	app.UsageText = app.Name + " <FOV><h|v> <aspect ratio> [new aspect ratio]"

	app.Action = func(c *cli.Context) error {
		var fov fov
		numArgs := len(c.Args())

		if numArgs < 2 {
			fmt.Fprintf(
				color.Output,
				color.HiRedString("Error:")+
					" Not enough arguments supplied; expected 0 or 1 arguments (got %d).\n"+
					"Usage: "+app.UsageText+"\n",
				numArgs)
			os.Exit(1)
		}

		// The first aspect ratio string passed to the command line (required)
		aspectRatioString := c.Args().Get(1)
		fov.aspectRatio = fractionToFloat(aspectRatioString)

		newAspectRatioString := c.Args().Get(2)

		if numArgs == 3 {
			// The second aspect ratio string passed to the command line (optional)
			fov.newAspectRatio = fractionToFloat(newAspectRatioString)
		}

		if strings.Count(c.Args().Get(0), "h") == 1 {
			// Horizontal FOV given, calculate the vertical FOV
			fov.horizontal, _ = strconv.ParseFloat(strings.TrimSuffix(c.Args().Get(0), "h"), 64)
			fov.vertical = math.Atan(math.Tan(fov.horizontal*math.Pi/360)*1/fov.aspectRatio) * 360 / math.Pi
			fov.newHorizontal = math.Atan(math.Tan(fov.horizontal*math.Pi/360)*fov.newAspectRatio/fov.aspectRatio) * 360 / math.Pi
			fov.newVertical = fov.vertical
		} else if strings.Count(c.Args().Get(0), "v") == 1 {
			// Vertical FOV given, calculate the horizontal FOV
			fov.vertical, _ = strconv.ParseFloat(strings.TrimSuffix(c.Args().Get(0), "v"), 64)
			fov.horizontal = math.Atan(math.Tan(fov.vertical*math.Pi/360)*fov.aspectRatio) * 360 / math.Pi
			fov.newHorizontal = math.Atan(math.Tan(fov.horizontal*math.Pi/360)*fov.newAspectRatio/fov.aspectRatio) * 360 / math.Pi
			fov.newVertical = fov.vertical
		} else {
			fmt.Fprintln(
				color.Output,
				color.HiRedString("Error:"),
				"Ambiguous FOV value given; the value must have",
				"a 'h' (horizontal) or 'v' (vertical) suffix.",
			)
			os.Exit(1)
		}

		switch numArgs {
		case 3:
			// An aspect ratio to convert is supplied
			fmt.Fprintf(
				color.Output,
				"\n                   Orig.\tConverted\n"+
					"  Horizontal FOV   %s\t%s\n"+
					"    Vertical FOV   %s\t%s\n"+
					"    Aspect ratio   %s\t\t%s\n",
				color.HiCyanString(degreeString(fov.horizontal)),
				color.HiYellowString(degreeString(fov.newHorizontal)),
				color.HiCyanString(degreeString(fov.vertical)),
				color.HiYellowString(degreeString(fov.newVertical)),
				color.HiCyanString(aspectRatioString),
				color.HiYellowString(newAspectRatioString),
			)
		case 2:
			// No aspect ratio to convert to is supplied
			fmt.Fprintf(
				color.Output,
				"\n  Horizontal FOV   %s\n"+
					"    Vertical FOV   %s\n"+
					"    Aspect ratio   %s\n",
				color.HiCyanString(degreeString(fov.horizontal)),
				color.HiCyanString(degreeString(fov.vertical)),
				color.HiCyanString(aspectRatioString),
			)
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
