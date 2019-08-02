package driver

import (
	"github.com/metalblueberry/go3d/vec2"
)

type World interface {
	Limits() (x int64, y int64)
	IsBlocked(x int64, y int64) bool
}

type Vehicle interface {
	MoveTo(Position vec2.T, Orientation vec2.T)
	GetPosition() (Position vec2.T, Orientation vec2.T)
}

type Driver struct {
	World   World
	Vehicle Vehicle
}

func NewDriver(Position vec2.T, Orientation Orientations, World World, Vehicle Vehicle) *Driver {
	Vehicle.MoveTo(Position, Orientation.Vector())
	return &Driver{
		World:   World,
		Vehicle: Vehicle,
	}
}

// Peek returns the next position when trying to move in a given grid.
func (d Driver) Peek(dir Direction) (vec2.T, vec2.T) {
	position, orientation := d.Vehicle.GetPosition()
	switch dir {
	case Forward:
		position = position.Add(orientation)
	case Backward:
		position = position.Sub(orientation)
	case Left:
		orientation = orientation.Rotate90DegLeft()
	case Right:
		orientation = orientation.Rotate90DegRight()
	}

	maxX, maxY := d.World.Limits()
	return wrap(position, maxX, maxY), orientation
}

func wrap(vector vec2.T, maxX, maxY int64) vec2.T {
	x := (int64(vector.Slice()[0]) + maxX) % maxX
	y := (int64(vector.Slice()[1]) + maxY) % maxY
	return vec2.T{float32(x), float32(y)}
}

func getCoords(vect vec2.T) (x int64, y int64) {
	slice := vect.Slice()
	return int64(slice[0]), int64(slice[1])
}

// Move tries to move the vehicle in the desired direction, returns false if failed.
func (d Driver) Move(dir Direction) bool {
	position, orientation := d.Peek(dir)
	x, y := getCoords(position)
	blocked := d.World.IsBlocked(x, y)
	if blocked {
		return false
	}

	d.Vehicle.MoveTo(position, orientation)
	return true
}

// ExecuteSequence will execute all comands and return the position of the last succesfully executed command,
// If the execution finishes before the sequence, the second parameters is false
func (d Driver) ExecuteSequence(sequence string) (int, bool) {
	for i, char := range sequence {
		ok := d.Move(Direction(char))
		if !ok {
			return i, false
		}
	}
	return len(sequence), true
}

func (d Driver) Position() (x int64, y int64) {
	position, _ := d.Vehicle.GetPosition()
	return getCoords(position)
}
