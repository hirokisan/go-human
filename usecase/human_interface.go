package usecase

import (
	"github.com/hirokisan/go-human/entity"
)

type Human interface {
	FullName() entity.Name
	Type() string
}
