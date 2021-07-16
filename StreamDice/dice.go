package main

import (
	"fmt"
	"math/rand"
)

type Dice struct {
	Edge int
}

//NewDice returns dice object.
func NewDice(edge int) *Dice {
	return &Dice{
		Edge: edge,
	}
}

//Roll rolls a dice and return roll-result.
func (d *Dice) Roll() int {
	return (1 + rand.Intn(d.Edge))
}

//RollDicePool rolls dicepool and prints results.
func (d *Dice) RollDicePool(numThrows int) {
	for i := 1; i <= numThrows; i++ {
		fmt.Println("Roll d", d.Edge, ": ", d.Roll())
	}
}
