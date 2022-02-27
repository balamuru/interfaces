package main

import (
	"strings"

	"github.com/balamuru/interfaces/animal"
)

type BoringAnimal interface {
	Move() string
	Eat() string
}

type Duck string

func (d Duck) Move() string {
	return "waddle"
}

func (d Duck) Eat() string {
	return "peck"
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

func (h Horse) Dance() string {
	return strings.ToUpper(h.walkBehavior+"-"+h.walkBehavior)
}

func main() {

	//define instance of a struct and exercise functions on it
	horsie := Horse{walkBehavior: "clip-clop", eatBehavior: "chomp-chomp"}
	println(horsie.walkBehavior)
	println(horsie.eatBehavior)
	println()

	//impl interface from aother package
	var animal animal.Animal = horsie
	println(animal.Move())
	println(animal.Eat())
	println(animal.Dance())
	println()


	//impl interface defined in current pkg
	var animal2 BoringAnimal = horsie
	println(animal2.Move())
	println(animal2.Eat())
	// animal2.Dance() - not valid as Animal2 does not implement Dance()
	println()

	//try this with a duck
	var animal3 BoringAnimal = Duck("donald")
	duck, ok := animal3.(Duck)
	println(duck)
	println(ok)
	println(animal3.Move())
	println(animal3.Eat())
}
