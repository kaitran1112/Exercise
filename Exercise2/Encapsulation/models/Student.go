package models

import "fmt"

type Student struct {
	name  string
	class string
}

func (s *Student) Input(a string, b string) {
	s.name = a
	s.class = b
}
func (s Student) Introduce() {
	fmt.Printf("Hello! I'm %v and in %v\n", s.name, s.class)
}
func (s *Student) changeClass(c string) {
	s.class = c
}
func (s *Student) ActionChangeClass(c string) {
	s.changeClass(c)
}
