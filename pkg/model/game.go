package model

import (
	"bytes"
	"fmt"
	"image/color"
	"time"

	"github.com/codingconcepts/snek-ops/pkg/constants"
	"github.com/codingconcepts/snek-ops/pkg/entity"
	"github.com/codingconcepts/snek-ops/pkg/game"
	"github.com/codingconcepts/snek-ops/pkg/math"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	world      *game.World
	lastUpdate time.Time
	gameOver   bool
	faceSource *text.GoTextFaceSource
}

func (g *Game) Update() error {
	if g.gameOver {
		return nil
	}

	snakeRaw, ok := g.world.FirstEntity("snake")
	if !ok {
		return fmt.Errorf("missing snake entity")
	}
	snake := snakeRaw.(*entity.Snake)

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		snake.SetDirection(math.DirUp)
	} else if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		snake.SetDirection(math.DirDown)
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		snake.SetDirection(math.DirLeft)
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		snake.SetDirection(math.DirRight)
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		var err error
		if g, err = NewGame(); err != nil {
			return fmt.Errorf("resetting game: %w", err)
		}
	}

	if time.Since(g.lastUpdate) < constants.GameSpeed {
		return nil
	}
	g.lastUpdate = time.Now()

	for _, entity := range g.world.Entities() {
		if entity.Update(g.world) {
			g.gameOver = true
			return nil
		}
	}

	return nil
}

func NewGame() (*Game, error) {
	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.MPlus1pRegular_ttf))
	if err != nil {
		return nil, fmt.Errorf("error loading font: %v", err)
	}

	world := game.NewWorld()

	// Add snake.
	world.AddEntity(entity.NewSnake(math.Point{
		X: constants.ScreenWidth / constants.GridSize / 2,
		Y: constants.ScreenHeight / constants.GridSize / 2,
	}, math.DirRight))

	// Add food.
	world.AddEntity(entity.NewFood(color.RGBA{255, 0, 0, 255}))

	g := Game{
		world:      world,
		faceSource: s,
	}

	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight)
	ebiten.SetWindowTitle("Snakeroach")

	return &g, nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entity := range g.world.Entities() {
		entity.Draw(screen)
	}

	if g.gameOver {
		face := &text.GoTextFace{
			Source: g.faceSource,
			Size:   48,
		}

		t := "Game Over"
		h, w := text.Measure(t, face, face.Size)

		opt := &text.DrawOptions{}
		opt.GeoM.Translate(constants.ScreenWidth/2-w/2, constants.ScreenHeight/2-h/2)
		opt.ColorScale.ScaleWithColor(color.White)

		text.Draw(screen, t, face, opt)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return constants.ScreenWidth, outsideHeight
}
