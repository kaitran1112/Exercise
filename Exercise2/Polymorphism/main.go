package main

import "fmt"

type Shouter interface {
	shout()
}
type Human struct {
	sound string
}

type Animal struct {
	sound string
}

func (h Human) shout() {
	fmt.Println(h.sound)
}
func (a Animal) shout() {
	fmt.Println(a.sound)
}
func shouting(shouter Shouter) {
	shouter.shout()
}
func main() {
	human := Human{"hahhahaha"}
	animal := Animal{"ececece"}
	human.shout()
	animal.shout()
	shouting(human)
	shouting(animal)
}

// hahhahaha
// ececece
// hahhahaha
// ececece
