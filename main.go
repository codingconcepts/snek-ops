package main

import (
	"log"

	"github.com/codingconcepts/snek-ops/pkg/model"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g, err := model.NewGame()
	if err != nil {
		log.Fatalf("error creating new game: %v", err)
	}

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
