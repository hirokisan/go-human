package usecase

import (
	"github.com/hirokisan/go-human/entity"
)

// Adult is implement of Human interface
type Adult entity.Human

// CreateAdult
func CreateAdult() Human {
	return &Adult{
		ID:        1,
		FirstName: entity.Name("James"),
		LastName:  entity.Name("Steve"),
	}
}

// FullName
func (c *Adult) FullName() entity.Name {
	return c.FirstName + " " + c.LastName
}

// Type
func (c *Adult) Type() string {
	return "Adult"
}
