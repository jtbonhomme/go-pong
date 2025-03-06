package pong

import (
	"github.com/hajimehoshi/ebiten"
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
	g.ball.Img, _ = ebiten.NewImage(int(g.ball.Radius*2), int(g.ball.Radius*2), ebiten.FilterDefault)
	g.player1.Img, _ = ebiten.NewImage(g.player1.Width, g.player1.Height, ebiten.FilterDefault)
	g.player2.Img, _ = ebiten.NewImage(g.player2.Width, g.player2.Height, ebiten.FilterDefault)

	InitFonts()
}

func (g *Game) reset(screen *ebiten.Image, state GameState) {
	w, _ := screen.Size()
	g.state = state
	g.rally = 0
	g.level = 0
	if state == StartState {
		g.player1.Score = 0
		g.player2.Score = 0
	}
	g.player1.Position = Position{
		X: InitPaddleShift, Y: GetCenter(screen).Y}
	g.player2.Position = Position{
		X: float32(w - InitPaddleShift - InitPaddleWidth), Y: GetCenter(screen).Y}
	g.ball.Position = GetCenter(screen)
	g.ball.Velocity = Velocity{X: initBallVelocity, Y: initBallVelocity}
}

// Layout sets the screen layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return windowWidth, windowHeight
}
