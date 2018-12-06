package main

import (
	"fmt"
)

type eat interface {
	eat()
}
type aaa func()

func (op *aaa) eat() {
	fmt.Println("dog eat feels good")
}

///////////////////////////////////////////////
func dog() {
	fmt.Println("I'm a dog")
}
///////////////////////////////////////////////

func feelsGood(a eat) {
	a.eat()
}

func main() {
	b := aaa(dog)
	feelsGood(&b)
	b()
}