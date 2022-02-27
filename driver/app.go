package main
import (
	"fmt"
	"github.com/balamuru/interfaces/animal"
)




type Animal2 interface {
	Move() string
	Eat() string

}

type Horse struct {
	walkBehavior string
	eatBehavior string
}

func (h Horse) Move() string{
	return h.walkBehavior
}

func (h Horse) Eat() string{
	return h.eatBehavior
}

func main() {

	fmt.Println("Hello World")
	horsie := Horse{walkBehavior: "clip-clop", eatBehavior: "chomp-chomp"}
	println(horsie.walkBehavior)
	println(horsie.eatBehavior)

	//impl interface from aother
	var animal animal.Animal = horsie 
	println(animal.Move())
	println(animal.Eat())


	var animal2 Animal2 = horsie 
	println(animal2.Move())
	println(animal2.Eat())
}