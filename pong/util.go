package pong

import (
	"image/color"
)

// Position is a set of coordinates in 2-D plan
type Position struct {
	X, Y float32
}

// Velocity is a set of coordinates in 2-D plan
type Velocity struct {
	X, Y float32
}

// GetCenter returns the center position on screen
func GetCenter(w, h int) Position {
	return Position{
		X: float32(w / 2),
		Y: float32(h / 2),
	}
}

// GameState is an enum that represents all possible game states
type GameState byte

const (
	StartState GameState = iota
	PlayState
	GameOverState
)

var (
	BgColor  = color.Black
	ObjColor = color.RGBA{120, 226, 160, 255}
)
