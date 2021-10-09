package main

import (
	"errors"
	"fmt"
)

type stack []interface{}

//Len returns the lenght of current stack.
func (stack stack) Len() int {
	return len(stack)
}

//Push adds element at the head of stack.
func (stack *stack) Push(x interface{}) {
	*stack = append(*stack, x)
}

//Top returns the top-element of current stack.
func (stack stack) Top() (interface{}, error) {
	if len(stack) == 0 {
		return nil, errors.New("can't Top() an empty stack")
	}
	return stack[len(stack)-1], nil
}

//Pop deletes top-element from stack.
func (stack *stack) Pop() (interface{}, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return nil, errors.New("can't Pop() an empty stack")
	}
	x := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return x, nil
}

func main() {
	var haystack stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)
	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
