package engine

func ResolveCollisions(parlicles []Particle) {
	for i := 0; i < len(parlicles); i++ {
		for j := i + 1; j < len(parlicles); j++ {
			ResolveParticleCollision(&parlicles[i], &parlicles[j])
		}
	}
}

func ResolveParticleCollision(p1, p2 *Particle) {
	delta := p2.Position.Sub(p1.Position)
	dist := delta.Length()
	minDist := p1.Radius + p2.Radius

	if dist < minDist {
		correction := delta.Normalize().Mul(minDist - dist)
		massRatio1 := p1.Radius / (p1.Radius + p2.Radius)
		massRatio2 := p2.Radius / (p1.Radius + p2.Radius)
		p1.Position = p1.Position.Sub(correction.Mul(massRatio2))
		p2.Position = p2.Position.Add(correction.Mul(massRatio1))
	}
}

func ResolveWallCollisions(parlicles []Particle, width, height float64) {
	for i := range parlicles {
		p := &parlicles[i]

		if p.Position.X-p.Radius < 0 {
			p.Position.X = p.Radius
		} else if p.Position.X+p.Radius > width {
			p.Position.X = width - p.Radius
		}

		if p.Position.Y-p.Radius < 0 {
			p.Position.Y = p.Radius
		} else if p.Position.Y+p.Radius > height {
			p.Position.Y = height - p.Radius
		}
	}
}
