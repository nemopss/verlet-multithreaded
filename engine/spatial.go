package engine

type Grid struct {
	CellSize float64
	Cells    map[int]map[int][]*Particle
}

func NewGrid(cellSize float64) *Grid {
	return &Grid{
		CellSize: cellSize,
		Cells:    make(map[int]map[int][]*Particle),
	}
}

func (g *Grid) Clear() {
	g.Cells = make(map[int]map[int][]*Particle)
}

func (g *Grid) GetCellCoords(pos Vec2) (int, int) {
	return int(pos.X / g.CellSize), int(pos.Y / g.CellSize)
}

func (g *Grid) Add(p *Particle) {
	cellX, cellY := g.GetCellCoords(p.Position)
	if _, exists := g.Cells[cellX]; !exists {
		g.Cells[cellX] = make(map[int][]*Particle)
	}
	g.Cells[cellX][cellY] = append(g.Cells[cellX][cellY], p)
}

func (g *Grid) GetNearby(p *Particle) []*Particle {
	cellX, cellY := g.GetCellCoords(p.Position)
	nearby := []*Particle{}
	for x := cellX - 1; x <= cellX+1; x++ {
		for y := cellY - 1; y <= cellY+1; y++ {
			if cells, exists := g.Cells[x][y]; exists {
				nearby = append(nearby, cells...)
			}
		}
	}
	return nearby
}
