package main

import "fmt"

type Animal struct {
	name   string
	sounds string
	legs   int
}

func (a Animal) sound() {
	fmt.Printf("%v sounds %v\n", a.name, a.sounds)
}
func (a Animal) countLeg() {
	fmt.Printf("%v has %d legs\n", a.name, a.legs)
}

type Cat struct {
	Animal
}
type Dog struct {
	Animal
}

func main() {
	cat1 := Cat{Animal{"Gao", "meow meow", 4}}
	cat1.sound()
	cat1.countLeg()
	dog1 := Dog{Animal{"Bull", "Gau Gau", 3}}
	dog1.sound()
	dog1.countLeg()
}

// Gao sounds meow meow
// Gao has 4 legs
// Bull sounds Gau Gau
// Bull has 3 legs
