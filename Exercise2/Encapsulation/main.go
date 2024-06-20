package main

import "Encapsulation/models"

func main() {
	var student1 models.Student
	student1.Input("kai", "K66CC")
	student1.Introduce()
	student1.ActionChangeClass("K66CB")
	student1.Introduce()
}

// Hello! I'm kai and in K66CC
// Hello! I'm kai and in K66CB
