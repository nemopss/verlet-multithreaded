package visualization

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/nemopss/verlet-multithreaded/engine"
)

// Renderer отвечает за визуализацию
type Renderer struct {
	Engine *engine.Engine
}

// Draw отрисовывает все частицы
func (r *Renderer) Draw(screen *ebiten.Image) {
	for _, p := range r.Engine.Particles {
		vector.DrawFilledCircle(screen, float32(p.Position.X), float32(p.Position.Y), float32(p.Radius), rgbaToColor(p.Color), true)
	}
}

// rgbaToColor преобразует RGBA в формат ebiten
func rgbaToColor(colors [4]uint8) color.RGBA {

	return color.RGBA{
		R: colors[0],
		G: colors[1],
		B: colors[2],
		A: colors[3],
	}
}
