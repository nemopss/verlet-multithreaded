package engine

import "sync"

type Engine struct {
	Particles []Particle
	Gravity   Vec2
	SubSteps  int
}

func NewEngine(gravity Vec2, subSteps int) *Engine {
	return &Engine{
		Particles: make([]Particle, 0),
		Gravity:   gravity,
		SubSteps:  subSteps,
	}
}

func (e *Engine) AddParticle(p Particle) {
	e.Particles = append(e.Particles, p)
}

func (e *Engine) Update(dt, width, height float64) {
	subDt := dt / float64(e.SubSteps)

	var wg sync.WaitGroup
	for step := 0; step < e.SubSteps; step++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range e.Particles {
				e.Particles[i].Acceleration = e.Gravity
				e.Particles[i].UpdatePosition(subDt)
			}
			ResolveCollisions(e.Particles)
			ResolveWallCollisions(e.Particles, width, height)
		}()
		wg.Wait()
	}
}
