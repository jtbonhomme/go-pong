package pong

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Update updates the game state
func (g *Game) Update() error {
	switch g.state {
	case StartState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = PlayState
		}

	case PlayState:
		g.player1.Update(g.height)
		g.player2.Update(g.height)

		xV := g.ball.Velocity.X
		g.ball.Update(g.player1, g.player2, g.height)
		// rally count
		if xV*g.ball.Velocity.X < 0 {
			// score up when ball touches human player's paddle
			if g.ball.Position.X < float32(g.width/2) {
				g.player1.Score++
			}

			g.rally++

			// spice things up
			if (g.rally)%speedUpdateCount == 0 {
				g.level++
				g.ball.Velocity.X += speedIncrement
				g.ball.Velocity.Y += speedIncrement
				g.player1.Speed += speedIncrement
				g.player2.Speed += speedIncrement
			}
		}

		if g.ball.Position.X < 0 {
			g.player2.Score++
			g.reset(StartState)
		} else if g.ball.Position.X > float32(g.width) {
			g.player1.Score++
			g.reset(StartState)
		}

		if g.player1.Score == g.maxScore || g.player2.Score == g.maxScore {
			g.state = GameOverState
		}

	case GameOverState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset(StartState)
		}
	}

	return nil
}
