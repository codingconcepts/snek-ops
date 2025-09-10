package math

import (
	"math/rand/v2"
	"slices"

	"github.com/codingconcepts/snek-ops/pkg/constants"
)

var (
	DirUp    = Point{X: 0, Y: -1}
	DirDown  = Point{X: 0, Y: 1}
	DirLeft  = Point{X: -1, Y: 0}
	DirRight = Point{X: 1, Y: 0}
)

type Point struct {
	X int
	Y int
}

func (p Point) Equals(other Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Point) Add(other Point) Point {
	return Point{
		X: p.X + other.X,
		Y: p.Y + other.Y,
	}
}

func (p Point) Collided(others []Point) bool {
	if p.X < 0 || p.Y < 0 {
		return true
	}

	if p.X >= constants.ScreenWidth/constants.GridSize || p.Y >= constants.ScreenHeight/constants.GridSize {
		return true
	}

	return slices.Contains(others, p)
}

func RandomPoint() Point {
	return Point{
		X: rand.IntN(constants.ScreenWidth / constants.GridSize),
		Y: rand.IntN(constants.ScreenHeight / constants.GridSize),
	}
}
