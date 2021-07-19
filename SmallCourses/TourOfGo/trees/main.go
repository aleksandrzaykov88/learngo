package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

//CountNode returns amount of tree leafs.
func CountNode(t *tree.Tree) int {
	if t == nil {
		return 0
	}
	return 1 + CountNode(t.Left) + CountNode(t.Right)
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 != nil {
		ch1 := make(chan int)
		ch2 := make(chan int)
		go Walk(t1, ch1)
		go Walk(t2, ch2)
		if CountNode(t1) == CountNode(t2) {
			for i := 0; i < CountNode(t1); i++ {
				if <-ch1 != <-ch2 {
					return false
				}
				return true
			}
		}
	}
	return false
}

func main() {
	ch := make(chan int)
	t := tree.New(2)
	go Walk(t, ch)
	for i := 0; i < CountNode(t); i++ {
		fmt.Println(<-ch)
	}
	t1, t2 := tree.New(2), tree.New(2)
	fmt.Println(Same(t1, t2))
}
