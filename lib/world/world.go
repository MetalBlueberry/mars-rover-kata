package world

type Grid struct {
	Width  int64
	Height int64
	tiles  []*Tile
}

func New(Width, Height int64) *Grid {
	g := &Grid{
		Width:  Width,
		Height: Height,
		tiles:  make([]*Tile, Width*Height),
	}
	for i, _ := range g.tiles {
		g.tiles[i] = &Tile{}
	}
	return g
}

func (g *Grid) Get(x, y int64) *Tile {
	return g.tiles[x+g.Width*y]
}

func (g *Grid) Limits() (int64, int64) {
	return g.Width, g.Height
}

func (g *Grid) IsBlocked(x, y int64) bool {
	return g.Get(x, y).Blocked
}

type Tile struct {
	Blocked bool
}
