package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Fov struct {
	horizontal  float64 // The "old" horizontal field of view
	vertical    float64 // The "old" vertical field of view
	aspectRatio float64 // The "old" aspect ratio

	// The following fields are only used if 3 arguments were given on the command line
	newHorizontal  float64 // The "new" horizontal field of view
	newVertical    float64 // The "new" vertical field of view
	newAspectRatio float64 // The "new" aspect ratio to scale the FOV to
}

// Converts a fraction in 'x:y' or 'x/y' notation to a floating-point value.
// Will return an error and exit if the input is not a fraction.
func fractionToFloat(fraction string) float64 {
	var terms []string

	if strings.Count(fraction, ":") == 1 {
		terms = strings.Split(fraction, ":")
	} else if strings.Count(fraction, "/") == 1 {
		terms = strings.Split(fraction, "/")
	} else {
		fmt.Println("Error: Bad format for aspect ratio; the ratio must have a 'x:y' or 'x/y' format.")
		os.Exit(1)
	}

	numerator, _ := strconv.ParseFloat(terms[0], 64)
	denominator, _ := strconv.ParseFloat(terms[1], 64)

	return numerator / denominator
}

func main() {
	app := cli.NewApp()
	app.Name = "fov"
	app.Version = "0.0.1"
	app.Usage = "Calculate horizontal or vertical FOV values for a given aspect ratio"
	app.UsageText = app.Name + " <FOV><h|v> <aspect ratio> [new aspect ratio]"

	app.Action = func(c *cli.Context) error {
		var fov Fov
		numArgs := len(c.Args())

		if numArgs < 2 {
			fmt.Printf("Error: Not enough arguments supplied; expected 2 or 3 arguments (got %d).\n"+
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
			fov.newHorizontal = math.Atan(math.Tan(fov.horizontal*math.Pi/360)*fov.newAspectRatio/fov.aspectRatio) * 360 / math.Pi
		} else if strings.Count(c.Args().Get(0), "v") == 1 {
			// Vertical FOV given, calculate the horizontal FOV
			fov.vertical, _ = strconv.ParseFloat(strings.TrimSuffix(c.Args().Get(0), "v"), 64)
			fov.newVertical = math.Atan(math.Tan(fov.vertical*math.Pi/360)*fov.newAspectRatio/fov.aspectRatio) * 360 / math.Pi
		} else {
			fmt.Println("Error: Ambiguous FOV value given; the value must have" +
				" a 'h' (horizontal) or 'v' (vertical) suffix.")
			os.Exit(1)
		}

		switch numArgs {
		case 3:
			// An aspect ratio to convert is supplied
			fmt.Printf("\t\tOrig.\tConverted\n"+
				"Horizontal FOV\t%.2f°\t%.2f°\n"+
				"Vertical FOV\t%.2f°\t%.2f°\n"+
				"Aspect ratio\t%s\t%s\n",
				fov.horizontal,
				fov.newHorizontal,
				fov.vertical,
				fov.newVertical,
				aspectRatioString,
				newAspectRatioString)
		case 2:
			// No aspect ratio to convert to is supplied
			fmt.Printf("Horizontal FOV\t%.2f°\n"+
				"Vertical FOV\t%.2f°\n"+
				"Aspect ratio\t%s\n",
				fov.horizontal,
				fov.vertical,
				aspectRatioString)
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
