package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nemopss/verlet-multithreaded/engine"
	"github.com/nemopss/verlet-multithreaded/visualization"
	"log"
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
	e := engine.NewEngine(engine.Vec2{X: 0, Y: 1000}, 8, 30)
	r := &visualization.Renderer{Engine: e}
	inputHandler := &visualization.InputHandler{Engine: e}
	game := &Game{Renderer: r, InputHandler: inputHandler}

	ebiten.SetWindowSize(1000, 1000)
	ebiten.SetWindowTitle("Verlet Physics")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
