package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/jtbonhomme/go-pong/pong"
)

// Game is the structure of the game state
type Game struct {
	state    pong.GameState
	ball     *pong.Ball
	player1  *pong.Paddle
	player2  *pong.Paddle
	rally    int
	level    int
	maxScore int
}

const (
	initBallVelocity = 5.0
	initPaddleSpeed  = 10.0
	speedUpdateCount = 6
	speedIncrement   = 0.5
)

const (
	windowWidth  = 800
	windowHeight = 600
)

// NewGame creates an initializes a new game
func NewGame() *Game {
	g := &Game{}
	g.init()
	return g
}

func (g *Game) init() {
	g.state = pong.StartState
	g.maxScore = 11

	g.player1 = &pong.Paddle{
		Position: pong.Position{
			X: pong.InitPaddleShift,
			Y: float32(windowHeight / 2)},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  pong.InitPaddleWidth,
		Height: pong.InitPaddleHeight,
		Color:  pong.ObjColor,
		Up:     ebiten.KeyUp,
		Down:   ebiten.KeyDown,
	}
	g.player2 = &pong.Paddle{
		Position: pong.Position{
			X: windowWidth - pong.InitPaddleShift - pong.InitPaddleWidth,
			Y: float32(windowHeight / 2)},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  pong.InitPaddleWidth,
		Height: pong.InitPaddleHeight,
		Color:  pong.ObjColor,
		Up:     ebiten.KeyW,
		Down:   ebiten.KeyS,
	}
	g.ball = &pong.Ball{
		Position: pong.Position{
			X: float32(windowWidth / 2),
			Y: float32(windowHeight / 2)},
		Radius:   pong.InitBallRadius,
		Color:    pong.ObjColor,
		Velocity: pong.Velocity{X: initBallVelocity, Y: initBallVelocity},
	}
	g.level = 0
	g.ball.Img, _ = ebiten.NewImage(int(g.ball.Radius*2), int(g.ball.Radius*2), ebiten.FilterDefault)
	g.player1.Img, _ = ebiten.NewImage(g.player1.Width, g.player1.Height, ebiten.FilterDefault)
	g.player2.Img, _ = ebiten.NewImage(g.player2.Width, g.player2.Height, ebiten.FilterDefault)

	pong.InitFonts()
}

func (g *Game) reset(screen *ebiten.Image, state pong.GameState) {
	w, _ := screen.Size()
	g.state = state
	g.rally = 0
	g.level = 0
	if state == pong.StartState {
		g.player1.Score = 0
		g.player2.Score = 0
	}
	g.player1.Position = pong.Position{
		X: pong.InitPaddleShift, Y: pong.GetCenter(screen).Y}
	g.player2.Position = pong.Position{
		X: float32(w - pong.InitPaddleShift - pong.InitPaddleWidth), Y: pong.GetCenter(screen).Y}
	g.ball.Position = pong.GetCenter(screen)
	g.ball.Velocity = pong.Velocity{X: initBallVelocity, Y: initBallVelocity}
}

// Update updates the game state
func (g *Game) Update(screen *ebiten.Image) error {
	switch g.state {
	case pong.StartState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.state = pong.PlayState
		}

	case pong.PlayState:
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
			g.reset(screen, pong.StartState)
		} else if g.ball.Position.X > float32(w) {
			g.player1.Score++
			g.reset(screen, pong.StartState)
		}

		if g.player1.Score == g.maxScore || g.player2.Score == g.maxScore {
			g.state = pong.GameOverState
		}

	case pong.GameOverState:
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.reset(screen, pong.StartState)
		}
	}

	g.Draw(screen)

	return nil
}

// Draw updates the game screen elements drawn
func (g *Game) Draw(screen *ebiten.Image) error {
	screen.Fill(pong.BgColor)

	pong.DrawCaption(g.state, pong.ObjColor, screen)
	pong.DrawBigText(g.state, pong.ObjColor, screen)
	g.player1.Draw(screen, pong.ArcadeFont)
	g.player2.Draw(screen, pong.ArcadeFont)
	g.ball.Draw(screen)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))

	return nil
}

// Layout sets the screen layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}

func main() {
	g := NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
