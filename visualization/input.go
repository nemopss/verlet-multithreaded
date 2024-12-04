package visualization

import (
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/nemopss/verlet-multithreaded/engine"
)

type InputHandler struct {
	Engine *engine.Engine
}

func (ih *InputHandler) HandleInput() {
	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		ih.Engine.Grid.CellSize += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyE) && ih.Engine.Grid.CellSize > 1 {
		ih.Engine.Grid.CellSize -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyZ) {
		ih.Engine.SubSteps += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyX) && ih.Engine.SubSteps > 1 {
		ih.Engine.SubSteps -= 1
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		color := [4]uint8{
			uint8(rand.Intn(256)), // R
			uint8(rand.Intn(256)), // G
			uint8(rand.Intn(256)), // B
			255,                   // A
		}

		// Добавляем новую частицу
		ih.Engine.AddParticle(engine.Particle{
			Position:     engine.Vec2{X: float64(x), Y: float64(y)},
			PrevPosition: engine.Vec2{X: float64(x), Y: float64(y)}.Sub(engine.RandomDirectionDown().Mul(5)), // Начальное движение
			Radius:       2 + rand.Float64()*3,
			Color:        color,
		})
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
