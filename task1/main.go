package main

import (
	"fmt"
)

type Human struct {
	Gender string
	Age    int
}

func (h Human) Introduce() {
	fmt.Printf("Human is a %v, %v years old\n", h.Gender, h.Age)
}

type Action struct {
	Human
}

func (a Action) Scream() {
	fmt.Println("aaaaaaaa")
	return
}

func main() {
	andrew := Action{
		Human: Human{"male", 12},
	}
	andrew.Introduce()
	andrew.Scream()
}
