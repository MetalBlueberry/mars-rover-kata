package rover

import (
	"github.com/metalblueberry/go3d/vec2"
)

type World interface {
	Limits() (x int64, y int64)
}

type Rover struct {
	Position    vec2.T
	Orientation vec2.T
	World       World
}

func NewRover(w World) *Rover {
	return &Rover{
		Position:    vec2.T{},
		Orientation: vec2.UnitX,
		World:       w,
	}
}

// Peek returns the next position when trying to move in a given grid.
func (r *Rover) Peek(dir Direction) vec2.T {
	var pos vec2.T
	switch dir {
	case Forward:
		pos = vec2.Add(r.Position, r.Orientation)
	case Backward:
		pos = vec2.Sub(r.Position, r.Orientation)
	default:
		return r.Position
	}
	maxX, maxY := r.World.Limits()
	return wrap(pos, maxX, maxY)
}

func (r *Rover) wrap() {
	maxX, maxY := r.World.Limits()
	r.Position = wrap(r.Position, maxX, maxY)
}

func wrap(vector vec2.T, maxX, maxY int64) vec2.T {
	x := (int64(vector.Slice()[0]) + maxX) % maxX
	y := (int64(vector.Slice()[1]) + maxY) % maxY
	return vec2.T{float32(x), float32(y)}
}

func (r *Rover) Move(dir Direction) bool {
	defer r.wrap()

	switch dir {
	case Forward:
		r.Position = r.Position.Add(r.Orientation)
	case Backward:
		r.Position = r.Position.Sub(r.Orientation)
	case Left:
		r.Orientation = r.Orientation.Rotate90DegLeft()
	case Right:
		r.Orientation = r.Orientation.Rotate90DegRight()
	default:
		return false
	}
	return true
}

// ExecuteSequence will execute all comands and return the position of the last succesfully executed command,
// If the execution finishes before the sequence, the second parameters is false
func (r *Rover) ExecuteSequence(sequence string) (int, bool) {
	for i, char := range sequence {
		ok := r.Move(Direction(char))
		if !ok {
			return i, false
		}
	}
	return len(sequence), true
}
