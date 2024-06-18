package main

import "fmt"

type Shouter interface {
	shout()
}
type ThingShouter struct {
	Shouter
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
	h := ThingShouter{human}
	a := ThingShouter{animal}
	human.shout()
	animal.shout()
	shouting(h)
	shouting(a)
}

// hahhahaha
// ececece
// hahhahaha
// ececece
