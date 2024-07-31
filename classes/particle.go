package classes

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/wrongheaven/elastic-collisions/consts"
	"github.com/wrongheaven/elastic-collisions/utils"
)

type Particle struct {
	Pos  utils.Vector2
	Vel  utils.Vector2
	Acc  utils.Vector2
	Mass float64
	R    float64
}

func (p Particle) New(x float64, y float64) Particle {
	p.Pos = utils.Vector2{}.New(x, y)
	p.Vel = utils.Vector2{}.Random()
	p.Vel.Mult(rand.Float64() * 4)
	p.Acc = utils.Vector2{}.Null()
	p.Mass = rand.Float64() * 6
	p.R = math.Sqrt(p.Mass) * 20

	return p
}

func (p *Particle) ApplyForce(force utils.Vector2) {
	force.Div(p.Mass)
	p.Acc.Add(force)
}

func (p *Particle) Update() {
	p.Vel.Add(p.Acc)
	p.Pos.Add(p.Vel)
	p.Acc.Mult(0)
}

func (p *Particle) BounceEdges() {
	if p.Pos.X > consts.WINDOW_WIDTH-p.R {
		p.Pos.X = consts.WINDOW_WIDTH - p.R
		p.Vel.X *= -1
	} else if p.Pos.X < p.R {
		p.Pos.X = p.R
		p.Vel.X *= -1
	}

	if p.Pos.Y > consts.WINDOW_HEIGHT-p.R {
		p.Pos.Y = consts.WINDOW_HEIGHT - p.R
		p.Vel.Y *= -1
	} else if p.Pos.Y < p.R {
		p.Pos.Y = p.R
		p.Vel.Y *= -1
	}
}

func (p Particle) Draw(screen *ebiten.Image) {
	ebitenutil.DrawCircle(screen, p.Pos.X, p.Pos.Y, p.R, color.RGBA{255, 255, 255, 255})
}
