package driver

type World interface {
	Limits() (x int64, y int64)
	IsBlocked(x int64, y int64) bool
}

type Vehicle interface {
}

type Driver struct {
	World
}
