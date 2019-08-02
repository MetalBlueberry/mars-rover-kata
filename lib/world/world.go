package world

type Grid struct {
	Width  int64
	Height int64
	tiles  []*Tile
}

func New(Width, Height int64) *Grid {
	return &Grid{
		Width:  Width,
		Height: Height,
		tiles:  make([]*Tile, Width*Height),
	}
}

func (w *Grid) Get(x, y int64) *Tile {
	return w.tiles[x+w.Width*y]
}

func (w *Grid) Limits() (int64, int64) {
	return w.Width, w.Height
}

type Tile struct {
	Blocked bool
}
