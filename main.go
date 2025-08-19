package main

import "fmt"

type Human struct {
	Name string
}

func (h Human) SayMyName() {
	fmt.Println(h.Name)
}

type Action struct {
	Human
}

func main() {
	act := Action{
		Human: Human{
			Name: "Nikita",
		},
	}

	act.SayMyName()
}
