package classes

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
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
	p.Vel = utils.Vector2{}.Random().Mult(rand.Float64() * 4)
	p.Acc = utils.Vector2{}.Null()
	p.Mass = rand.Float64() * 6
	p.R = math.Sqrt(p.Mass) * 20

	return p
}

func (p *Particle) ApplyForce(force utils.Vector2) {
	p.Acc = p.Acc.Add(force.Div(p.Mass))
}

func (p *Particle) Update() {
	p.Vel = p.Vel.Add(p.Acc)
	p.Pos = p.Pos.Add(p.Vel)
	p.Acc = p.Acc.Mult(0)
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

func (p *Particle) CheckAndResolveCollision(other *Particle) {
	if p.Pos.Dist(other.Pos) < p.R+other.R {
		x1 := p.Pos
		x2 := other.Pos
		v1 := p.Vel
		v2 := other.Vel
		m1 := p.Mass
		m2 := other.Mass

		mSum := m1 + m2
		impact := x2.Sub(x1)
		vDiff := v2.Sub(v1)

		num := 2 * m2 * vDiff.Dot(impact)
		den := mSum * math.Pow(p.Pos.Dist(other.Pos), 2)
		nv1 := v1.Add(impact.Mult(num / den))

		impact = x1.Sub(x2)
		vDiff = v1.Sub(v2)

		num = 2 * m1 * vDiff.Dot(impact)
		nv2 := v2.Add(impact.Mult(num / den))

		p.Vel = nv1
		other.Vel = nv2
	}
}

func (p Particle) GetMomentum() utils.Vector2 {
	return p.Vel.Mult(p.Mass)
}
func (p Particle) GetKineticEnergy() float64 {
	mag := p.Vel.Mag()
	return 0.5 * p.Mass * mag * mag
}

func (p Particle) Show(screen *ebiten.Image) {
	vector.DrawFilledCircle(
		screen,
		float32(p.Pos.X),
		float32(p.Pos.Y),
		float32(p.R),
		color.RGBA{255, 255, 255, 255},
		true,
	)
}
