package main

import (
	"os"

	"github.com/metalblueberry/go3d/vec2"
	"github.com/metalblueberry/mars-rover-kata/lib/driver"
	"github.com/metalblueberry/mars-rover-kata/lib/rover"
	"github.com/metalblueberry/mars-rover-kata/lib/util"
	"github.com/metalblueberry/mars-rover-kata/lib/world"
)

func main() {
	World := world.New(100, 50)
	Rover := rover.New()
	Driver := driver.New(vec2.T{0, 0}, driver.East, World, Rover)
	util.DrawWorld(os.Stdout, World, Driver)
}
