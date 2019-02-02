package main

import (
	"fmt"

	"github.com/hirokisan/go-human/interface/factory"
)

func main() {
	child := factory.CreateHuman(factory.Child)
	fmt.Println(child.FullName())
	fmt.Println(child.Type())

	adult := factory.CreateHuman(factory.Adult)
	fmt.Println(adult.FullName())
	fmt.Println(adult.Type())
}
