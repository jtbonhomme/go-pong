package pong

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Draw updates the game screen elements drawn
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BgColor)

	DrawCaption(g.state, ObjColor, screen)
	DrawBigText(g.state, ObjColor, screen)
	g.player1.Draw(screen, ArcadeFont)
	g.player2.Draw(screen, ArcadeFont)
	g.ball.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
}
