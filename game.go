package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wrongheaven/elastic-collisions/classes"
	"github.com/wrongheaven/elastic-collisions/consts"
)

type Game struct {
	particleA classes.Particle
	particleB classes.Particle
}

func (g *Game) Layout(width int, height int) (int, int) {
	return consts.WINDOW_WIDTH, consts.WINDOW_HEIGHT
}

func (g *Game) Setup() {
	g.particleA = classes.Particle{}.New(320, 60)
	g.particleB = classes.Particle{}.New(320, 300)
}

func (g *Game) Update() error {
	g.particleA.Update()
	g.particleB.Update()

	g.particleA.BounceEdges()
	g.particleB.BounceEdges()

	g.particleA.CheckAndResolveCollision(&g.particleB)

	// fmt.Println(g.particleA.GetMomentum().Add(g.particleB.GetMomentum()))
	// fmt.Println(g.particleA.GetKineticEnergy() + g.particleB.GetKineticEnergy())

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{127, 127, 127, 255})

	g.particleA.Show(screen)
	g.particleB.Show(screen)
}
