package pong

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game is the structure of the game state
type Game struct {
	state    GameState
	ball     *Ball
	player1  *Paddle
	player2  *Paddle
	rally    int
	level    int
	maxScore int
	width    int
	height   int
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
	g := &Game{
		width:  windowWidth,
		height: windowHeight,
	}
	g.init()
	return g
}

func (g *Game) init() {
	g.state = StartState
	g.maxScore = 11

	g.player1 = &Paddle{
		Position: Position{
			X: InitPaddleShift,
			Y: float32(windowHeight / 2)},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  InitPaddleWidth,
		Height: InitPaddleHeight,
		Color:  ObjColor,
		Up:     ebiten.KeyUp,
		Down:   ebiten.KeyDown,
	}
	g.player2 = &Paddle{
		Position: Position{
			X: windowWidth - InitPaddleShift - InitPaddleWidth,
			Y: float32(windowHeight / 2)},
		Score:  0,
		Speed:  initPaddleSpeed,
		Width:  InitPaddleWidth,
		Height: InitPaddleHeight,
		Color:  ObjColor,
		Up:     ebiten.KeyW,
		Down:   ebiten.KeyS,
	}
	g.ball = &Ball{
		Position: Position{
			X: float32(windowWidth / 2),
			Y: float32(windowHeight / 2)},
		Radius:   InitBallRadius,
		Color:    ObjColor,
		Velocity: Velocity{X: initBallVelocity, Y: initBallVelocity},
	}
	g.level = 0
	g.ball.Img = ebiten.NewImage(int(g.ball.Radius*2), int(g.ball.Radius*2))
	g.player1.Img = ebiten.NewImage(g.player1.Width, g.player1.Height)
	g.player2.Img = ebiten.NewImage(g.player2.Width, g.player2.Height)

	InitFonts()
}

func (g *Game) reset(state GameState) {
	g.state = state
	g.rally = 0
	g.level = 0
	if state == StartState {
		g.player1.Score = 0
		g.player2.Score = 0
	}
	g.player1.Position = Position{
		X: InitPaddleShift, Y: GetCenter(g.width, g.height).Y}
	g.player2.Position = Position{
		X: float32(g.width - InitPaddleShift - InitPaddleWidth), Y: GetCenter(g.width, g.height).Y}
	g.ball.Position = GetCenter(g.width, g.height)
	g.ball.Velocity = Velocity{X: initBallVelocity, Y: initBallVelocity}
}

// Layout sets the screen layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}
