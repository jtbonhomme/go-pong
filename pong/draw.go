package pong

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Draw updates the game screen elements drawn
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(BgColor)
	g.drawVerticalDottedLine(screen, float32(g.width/2), 50, float32(g.height)-50, 15)

	DrawCaption(g.state, ObjColor, screen)
	DrawBigText(g.state, ObjColor, screen)
	g.player1.Draw(screen, ArcadeFont)
	g.player2.Draw(screen, ArcadeFont)
	g.ball.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) drawVerticalDottedLine(screen *ebiten.Image, x0, y0, h, dot float32) {
	var strokeWidth float32 = 1

	for y := y0; y < y0+h; y += 2 * dot {
		vector.StrokeLine(screen, x0, y, x0, y+dot, strokeWidth, color.RGBA{0x8b, 0x8d, 0x80, 0xff}, false)
	}
}
