package main

import (
	"fmt"
	"strings"
)

type Person struct{ name string }

type Swimmer struct{}

func (s *Swimmer) Swim(name string) {
	fmt.Println(strings.Join([]string{
		name,
		" is swimming",
	}, ""))
}

type IronMan struct {
	person  Person
	swimmer Swimmer
}

func (i *IronMan) Swim() {
	i.swimmer.Swim(i.person.name)
}

func main() {
	ironMan := IronMan{
		person:  Person{"Mariottide"},
		swimmer: Swimmer{},
	}

	ironMan.Swim()
}
