package entity

import (
	"image/color"

	"github.com/codingconcepts/snek-ops/pkg/constants"
	"github.com/codingconcepts/snek-ops/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Ensure our Food struct implements entity.
var _ Entity = (*Food)(nil)

type Food struct {
	position math.Point
	color.Color
}

func NewFood(c color.Color) *Food {
	return &Food{
		position: math.RandomPoint(),
		Color:    c,
	}
}

func (f *Food) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		float32(f.position.X*constants.GridSize),
		float32(f.position.Y*constants.GridSize),
		constants.GridSize,
		constants.GridSize,
		f.Color,
		true,
	)
}

func (f *Food) Respawn() {
	f.position = math.RandomPoint()
}

func (f *Food) Tag() string {
	return "food"
}

func (f *Food) Update(world WorldView) bool {
	return false
}
