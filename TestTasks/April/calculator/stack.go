package main

import "errors"

//Stack classic implementation.
type Stack []float64

func (s *Stack) Push(value float64) {
	*s = append(*s, value)
}

func (s *Stack) Pop() (float64, error) {
	temp, err := s.Peek()
	if err != nil {
		return 0.0, err
	}

	*s = (*s)[0 : len(*s)-1]

	return temp, nil
}

func (s *Stack) Peek() (float64, error) {
	if len(*s) == 0 {
		return 0.0, errors.New("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

func (s *Stack) Len() int {
	return len(*s)
}
