package rover

import (
	"github.com/metalblueberry/go3d/vec2"
)

type Rover struct {
	Position    vec2.T
	Orientation vec2.T
}

func NewRover() *Rover {
	return &Rover{
		Position:    vec2.T{},
		Orientation: vec2.UnitX,
	}
}

func (r *Rover) MoveTo(Position vec2.T, Orientation vec2.T) {
	r.Position = Position
	r.Orientation = Orientation
}
func (r *Rover) GetPosition() (Position vec2.T, Orientation vec2.T) {
	return r.Position, r.Orientation
}
