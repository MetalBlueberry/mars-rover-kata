package main

import (
	"fmt"

	"github.com/metalblueberry/mars-rover-kata/lib/rover"
)

func main() {
	dir := rover.Direction('X')
	fmt.Print(dir)
}
