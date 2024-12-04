package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nemopss/verlet-multithreaded/engine"
	"github.com/nemopss/verlet-multithreaded/visualization"
)

type Game struct {
	Renderer     *visualization.Renderer
	InputHandler *visualization.InputHandler
}

func (g *Game) Update() error {
	screenWidth, screenHeight := ebiten.WindowSize()
	g.InputHandler.HandleInput()
	g.Renderer.Engine.Update(1.0/60.0, float64(screenWidth), float64(screenHeight))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Renderer.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func main() {
	e := engine.NewEngine(engine.Vec2{X: 0, Y: 1000}, 4, 10)
	for i := 0; i < 90; i++ {
		for j := 0; j < 40; j++ {
			e.AddParticle(engine.Particle{
				Position:     engine.Vec2{X: float64(0 + i*10), Y: float64(0 + j*10)},
				PrevPosition: engine.Vec2{X: float64(0 + i*10), Y: float64(0 + j*10)},
				Acceleration: engine.Vec2{X: 0, Y: 0},
				Radius:       4,
				Color:        [4]uint8{255 - uint8(j*5), 255 - uint8(i*2), 0, 255},
			})
		}
	}

	r := &visualization.Renderer{Engine: e}
	inputHandler := &visualization.InputHandler{Engine: e}
	game := &Game{Renderer: r, InputHandler: inputHandler}
	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Verlet Physics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
