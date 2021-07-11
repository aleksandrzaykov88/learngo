package main

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Edge int
}

func NewDice(edge int) *Dice {
	return &Dice{
		Edge: edge,
	}
}

func (d *Dice) Roll() int {
	return (1 + rand.Intn(d.Edge))
}

func (d *Dice) RollDicePool(numThrows int) {
	for i := 1; i <= numThrows; i++ {
		fmt.Println("Roll d", d.Edge, ": ", d.Roll())
	}
}
