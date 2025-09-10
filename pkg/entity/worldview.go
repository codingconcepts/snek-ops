package entity

type WorldView interface {
	Entities() []Entity
	FindEntities(tag string) []Entity
}
