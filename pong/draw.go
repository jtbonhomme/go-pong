package pong

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Draw updates the game screen elements drawn
func (g *Game) Draw(screen *ebiten.Image) error {
	screen.Fill(BgColor)

	DrawCaption(g.state, ObjColor, screen)
	DrawBigText(g.state, ObjColor, screen)
	g.player1.Draw(screen, ArcadeFont)
	g.player2.Draw(screen, ArcadeFont)
	g.ball.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	return nil
}
