package utils

import (
	"math"
	"math/rand/v2"
)

type Vector2 struct {
	X, Y float64
}

func (v Vector2) New(x float64, y float64) Vector2 {
	return Vector2{X: x, Y: y}
}
func (v Vector2) Null() Vector2 {
	return Vector2{X: 0, Y: 0}
}
func (v Vector2) Random() Vector2 {
	angle := rand.Float64() * 360
	return v.FromDeg(angle)
}
func (v Vector2) FromDeg(angle float64) Vector2 {
	angle = (angle + 90) * 2 * math.Pi / 360
	return v.FromRad(angle)
}
func (v Vector2) FromRad(angle float64) Vector2 {
	x, y := math.Sincos(angle)
	return Vector2{X: x, Y: y}
}

func (v Vector2) Add(other Vector2) Vector2 {
	v.X += other.X
	v.Y += other.Y
	return v
}
func (v Vector2) Sub(other Vector2) Vector2 {
	v.X -= other.X
	v.Y -= other.Y
	return v
}
func (v Vector2) Mult(n float64) Vector2 {
	v.X *= n
	v.Y *= n
	return v
}
func (v Vector2) Div(n float64) Vector2 {
	v.X /= n
	v.Y /= n
	return v
}

func (v Vector2) Dot(other Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v Vector2) Mag() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vector2) Dist(other Vector2) float64 {
	dx := math.Abs(v.X - other.X)
	dy := math.Abs(v.Y - other.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
