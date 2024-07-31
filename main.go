package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/wrongheaven/elastic-collisions/consts"
)

func main() {
	ebiten.SetWindowSize(consts.WINDOW_WIDTH, consts.WINDOW_HEIGHT)
	ebiten.SetWindowTitle(consts.WINDOW_TITLE)
	ebiten.SetTPS(consts.FPS)

	game := Game{}
	game.Setup()

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
