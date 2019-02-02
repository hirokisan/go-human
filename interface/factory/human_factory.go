package factory

import (
	"github.com/hirokisan/go-human/usecase"
)

// HumanType
type HumanType string

const (
	// Child is one of HumanType
	Child = HumanType("child")
	// Adult is one of HumanType
	Adult = HumanType("adult")
)

// CreateHuman create implemented human
func CreateHuman(typ HumanType) usecase.Human {
	switch typ {
	case Child:
		return usecase.CreateChild()
	case Adult:
		return usecase.CreateAdult()
	}
	return nil
}
