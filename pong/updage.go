package pong

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
)

// Update updates the game state
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case StartState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = PlayState
		}

	case PlayState:
		w, _ := screen.Size()

		g.player1.Update(screen)
		g.player2.Update(screen)

		xV := g.ball.Velocity.X
		g.ball.Update(g.player1, g.player2, screen)
		// rally count
		if xV*g.ball.Velocity.X < 0 {
			// score up when ball touches human player's paddle
			if g.ball.Position.X < float32(w/2) {
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
			g.reset(screen, StartState)
		} else if g.ball.Position.X > float32(w) {
			g.player1.Score++
			g.reset(screen, StartState)
		}

		if g.player1.Score == g.maxScore || g.player2.Score == g.maxScore {
			g.state = GameOverState
		}

	case GameOverState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset(screen, StartState)
		}
	}

	g.Draw(screen)

	return nil
}
