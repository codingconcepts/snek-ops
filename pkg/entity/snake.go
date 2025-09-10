package entity

import (
	"image/color"

	"github.com/codingconcepts/snek-ops/pkg/constants"
	"github.com/codingconcepts/snek-ops/pkg/math"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

// Ensure our Snake struct implements entity.
var _ Entity = (*Snake)(nil)

type Snake struct {
	body      []math.Point
	direction math.Point
}

func NewSnake(start, direction math.Point) *Snake {
	return &Snake{
		body:      []math.Point{start},
		direction: direction,
	}
}

func (s *Snake) Draw(screen *ebiten.Image) {
	for _, p := range s.body {
		vector.DrawFilledRect(
			screen,
			float32(p.X*constants.GridSize),
			float32(p.Y*constants.GridSize),
			constants.GridSize,
			constants.GridSize,
			color.White,
			true,
		)
	}
}

func (s *Snake) SetDirection(d math.Point) {
	// Don't let the player reverse their direction (crashing into themselves).
	switch d {
	case math.DirDown:
		if s.direction == math.DirUp {
			return
		}
	case math.DirUp:
		if s.direction == math.DirDown {
			return
		}
	case math.DirRight:
		if s.direction == math.DirLeft {
			return
		}
	case math.DirLeft:
		if s.direction == math.DirRight {
			return
		}
	}

	s.direction = d
}

func (s *Snake) Tag() string {
	return "snake"
}

func (s *Snake) Update(world WorldView) bool {
	newHead := s.body[0].Add(s.direction)

	if newHead.Collided(s.body) {
		return true
	}

	var grow bool

	for _, entity := range world.FindEntities("food") {
		food := entity.(*Food)

		if newHead.Equals(food.position) {
			grow = true
			food.Respawn()
			break
		}
	}

	if grow {
		s.body = append([]math.Point{newHead}, s.body...)
	} else {
		s.body = append([]math.Point{newHead}, (s.body)[:len(s.body)-1]...)
	}

	return false
}
