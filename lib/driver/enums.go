package driver

import "github.com/metalblueberry/go3d/vec2"

// Direction represent any of the possible input actions for Mars Rover
type Direction byte

const (
	UnknownDirection Direction = 0
	Forward          Direction = 'F'
	Backward         Direction = 'B'
	Left             Direction = 'L'
	Right            Direction = 'R'
)

type Orientations byte

const (
	UnknownOrientation Orientations = 0
	North              Orientations = 'N'
	South              Orientations = 'S'
	East               Orientations = 'E'
	Weast              Orientations = 'W'
)

func (o Orientations) Vector() vec2.T {
	switch o {
	case North:
		return vec2.UnitY.Inverted()
	case South:
		return vec2.UnitY
	case East:
		return vec2.UnitX
	case Weast:
		return vec2.UnitX.Inverted()
	}
	panic("Invalid orientation")
}
