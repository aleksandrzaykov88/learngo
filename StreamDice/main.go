package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var d20 = NewDice(20)
	d20.RollDicePool(15)
}
