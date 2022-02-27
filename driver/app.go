package main

import (
	"github.com/balamuru/interfaces/animal"
)

type Animal2 interface {
	Move() string
	Eat() string
}

type Horse struct {
	walkBehavior string
	eatBehavior  string
}

func (h Horse) Move() string {
	return h.walkBehavior
}

func (h Horse) Eat() string {
	return h.eatBehavior
}

func main() {

	//define instance of a struct and exercise functions on it
	horsie := Horse{walkBehavior: "clip-clop", eatBehavior: "chomp-chomp"}
	println(horsie.walkBehavior)
	println(horsie.eatBehavior)

	//impl interface from aother package
	var animal animal.Animal = horsie
	println(animal.Move())
	println(animal.Eat())

	//impl interface defined in current pkg
	var animal2 Animal2 = horsie
	println(animal2.Move())
	println(animal2.Eat())
}
