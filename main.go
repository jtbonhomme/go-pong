package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/jtbonhomme/go-pong/pong"
)

func main() {
	g := pong.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
