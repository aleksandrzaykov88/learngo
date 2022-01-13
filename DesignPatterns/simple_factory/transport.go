package main

import "fmt"

type transport struct {
	name  string
	speed int
}

func (t *transport) setName(name string) {
	t.name = name
}

func (t *transport) getName() string {
	return t.name
}

func (t *transport) deliver() {
	fmt.Println("Cargo has been delivered by " + t.getName())
}
