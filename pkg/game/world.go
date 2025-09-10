package game

import (
	"github.com/codingconcepts/snek-ops/pkg/entity"
)

type World struct {
	entities []entity.Entity
}

func NewWorld() *World {
	return &World{
		entities: []entity.Entity{},
	}
}

func (w *World) AddEntity(e entity.Entity) {
	w.entities = append(w.entities, e)
}

func (w *World) Entities() []entity.Entity {
	return w.entities
}

func (w *World) FindEntities(tag string) []entity.Entity {
	var found []entity.Entity

	for _, e := range w.entities {
		if e.Tag() == tag {
			found = append(found, e)
		}
	}

	return found
}

func (w *World) FirstEntity(tag string) (entity.Entity, bool) {
	for _, e := range w.entities {
		if e.Tag() == tag {
			return e, true
		}
	}

	return nil, false
}
