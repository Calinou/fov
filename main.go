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

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func truncate(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	app := cli.NewApp()
	app.Name = "fov"
	app.Usage = "Calculate horizontal or vertical FOV values for a given aspect ratio"

	app.Action = func(c *cli.Context) error {
		var fov Fov

		if strings.Count(c.Args().Get(0), "h") == 1 {
			// Horizontal FOV given, calculate the vertical FOV
			fov.horizontal, _ = strconv.ParseFloat(strings.TrimSuffix(c.Args().Get(0), "h"), 64)
		} else if strings.Count(c.Args().Get(0), "v") == 1 {
			// Vertical FOV given, calculate the horizontal FOV
			fov.vertical, _ = strconv.ParseFloat(strings.TrimSuffix(c.Args().Get(0), "v"), 64)
		} else {
			fmt.Println("Error: Ambiguous FOV value given; the value must have" +
				" a 'h' (horizontal) or 'v' (vertical) suffix.")
		}

		// Aspect ratio string passed to the command line
		// We can't convert the aspect ratio from the struct, it would suffer from precision loss
		aspectRatioString := c.Args().Get(1)

		switch len(c.Args()) {
		case 3:
			newAspectRatioString := c.Args().Get(2)

			fmt.Printf("\t\tOrig.\tConverted\n"+
				"Horizontal FOV\t%.1f°\t%.1f°\n"+
				"Vertical FOV\t%.1f°\t%.1f°\n"+
				"Aspect ratio\t%s\t%s\n",
				fov.horizontal, fov.newHorizontal, fov.vertical, fov.newVertical, aspectRatioString, newAspectRatioString)
		case 2:
			fmt.Printf("Horizontal FOV\t%.1f°\n"+
				"Vertical FOV\t%.1f°\n"+
				"Aspect ratio\t%s\n",
				fov.horizontal, fov.vertical, aspectRatioString)
		default:
			fmt.Println("Error: Too many arguments, 2 or 3 arguments expected.")
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
