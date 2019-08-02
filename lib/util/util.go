package util

import (
	"fmt"
	"io"

	"github.com/metalblueberry/mars-rover-kata/lib/driver"
	"github.com/metalblueberry/mars-rover-kata/lib/world"
)

// DrawWorld writes an ascii representations of the Grid and the Driver to out io.Writer
func DrawWorld(out io.Writer, Grid *world.Grid, Driver *driver.Driver) {

	maxX, maxY := Grid.Limits()

	fmt.Fprint(out, " +")
	for x := int64(0); x < maxX; x++ {
		if x%5 == 0 {
			fmt.Fprint(out, "v")
		} else {
			fmt.Fprint(out, "_")
		}
	}
	fmt.Fprint(out, "+")
	fmt.Fprint(out, "\n")
	for y := int64(0); y < maxY; y++ {
		fmt.Fprint(out, "⎹ ")
		for x := int64(0); x < maxX; x++ {
			if i, j := Driver.Position(); x == i && y == j {
				fmt.Fprint(out, "O")
				continue
			}
			if Grid.IsBlocked(x, y) {
				fmt.Fprint(out, "X")
				continue
			}
			fmt.Fprint(out, "-")
		}
		fmt.Fprint(out, "⎸")
		fmt.Fprint(out, "\n")
	}
	fmt.Fprint(out, " +")
	for x := int64(0); x < maxX; x++ {
		if x%5 == 0 {
			fmt.Fprint(out, "^")
		} else {
			fmt.Fprint(out, "‾")
		}
	}
	fmt.Fprint(out, "+")
	fmt.Fprint(out, "\n")

}
