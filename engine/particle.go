package engine

type Particle struct {
	Position     Vec2
	PrevPosition Vec2
	Acceleration Vec2
	Radius       float64
	Color        [4]uint8
}

func (p *Particle) UpdatePosition(dt float64) {
	temp := p.Position
	p.Position = p.Position.Add(p.Position.Sub(p.PrevPosition)).Add(p.Acceleration.Mul(dt * dt))
	p.PrevPosition = temp
}
