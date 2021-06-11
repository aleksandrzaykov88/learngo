package main

import (
	"errors"
)

type stack []interface{}

func (stack stack) Len() int {
	return len(stack)
}

func (stack *stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

func (stack stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

func (stack *stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}
