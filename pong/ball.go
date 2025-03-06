package pong

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
)

// Ball is a pong ball
type Ball struct {
	Position
	Velocity
	Radius float32
	Color  color.Color
	Img    *ebiten.Image
}

const (
	InitBallRadius = 10.0
)

func setBallPixels(c color.Color, ballImg *ebiten.Image) {
	// TODO: set pixels for round effect
	ballImg.Fill(c)
}

func (b *Ball) Update(leftPaddle *Paddle, rightPaddle *Paddle, screen *ebiten.Image) {
	_, h := screen.Size()
	b.Position.X += b.Velocity.X
	b.Position.Y += b.Velocity.Y

	// bounce off edges when getting to top/bottom
	if b.Position.Y-b.Radius > float32(h) {
		b.Velocity.Y = -b.Velocity.Y
		b.Position.Y = float32(h) - b.Radius
	} else if b.Position.Y+b.Radius < 0 {
		b.Velocity.Y = -b.Velocity.Y
		b.Position.Y = b.Radius
	}

	// bounce off paddles
	if b.Position.X-b.Radius < leftPaddle.X+float32(leftPaddle.Width/2) &&
		b.Position.Y > leftPaddle.Y-float32(leftPaddle.Height/2) &&
		b.Position.Y < leftPaddle.Y+float32(leftPaddle.Height/2) {
		b.Velocity.X = -b.Velocity.X
		b.Position.X = leftPaddle.X + float32(leftPaddle.Width/2) + b.Radius
	} else if b.Position.X+b.Radius > rightPaddle.X-float32(rightPaddle.Width/2) &&
		b.Position.Y > rightPaddle.Y-float32(rightPaddle.Height/2) &&
		b.Position.Y < rightPaddle.Y+float32(rightPaddle.Height/2) {
		b.Velocity.X = -b.Velocity.X
		b.Position.X = rightPaddle.X - float32(rightPaddle.Width/2) - b.Radius
	}
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(b.Position.X), float64(b.Position.Y))
	setBallPixels(b.Color, b.Img)
	screen.DrawImage(b.Img, opts)
}
