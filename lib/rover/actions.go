package rover

// Direction represent any of the possible input actions for Mars Rover
type Direction byte

const (
	Unknown  Direction = 0
	Forward  Direction = 'F'
	Backward Direction = 'B'
	Left     Direction = 'L'
	Right    Direction = 'R'
)
