package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nemopss/verlet-multithreaded/engine"
	"github.com/nemopss/verlet-multithreaded/visualization"
	"log"
)

type Game struct {
	Renderer *visualization.Renderer
}

func (g *Game) Update() error {
	g.Renderer.Engine.Update(1.0 / 60.0)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Renderer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	e := engine.NewEngine(engine.Vec2{X: 0, Y: 1000}, 4)
	r := &visualization.Renderer{Engine: e}

	// Добавление тестовых частиц
	e.AddParticle(engine.Particle{
		Position:     engine.Vec2{X: 100, Y: 100},
		PrevPosition: engine.Vec2{X: 100, Y: 100},
		Radius:       5,
		Color:        [4]uint8{255, 255, 255, 255},
	})

	game := &Game{Renderer: r}

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Verlet Physics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
