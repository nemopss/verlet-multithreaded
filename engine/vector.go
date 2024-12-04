package engine

import (
	"math"
	"math/rand/v2"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{v.X + other.X, v.Y + other.Y}
}

func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{v.X - other.X, v.Y - other.Y}
}

func (v Vec2) Mul(scalar float64) Vec2 {
	return Vec2{v.X * scalar, v.Y * scalar}
}

func (v Vec2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Normalize() Vec2 {
	length := v.Length()
	if length == 0 {
		return Vec2{0, 0}
	}
	return Vec2{v.X / length, v.Y / length}
}

func (v Vec2) Distance(other Vec2) float64 {
	return v.Sub(other).Length()
}

func RandomDirectionDown() Vec2 {
	angle := math.Pi/4 + rand.Float64()*(math.Pi/2) // Угол в диапазоне [π/4, 3π/4]
	return Vec2{
		X: math.Cos(angle),
		Y: math.Sin(angle),
	}
}
