package usecase

import (
	"github.com/hirokisan/go-human/entity"
)

// Child is implement of Human interface
type Child entity.Human

// CreateChild
func CreateChild() Human {
	return &Child{
		ID:        1,
		FirstName: entity.Name("James"),
		LastName:  entity.Name("Steve"),
	}
}

// FullName
func (c *Child) FullName() entity.Name {
	return c.FirstName + " " + c.LastName
}

// Type
func (c *Child) Type() string {
	return "Child"
}
