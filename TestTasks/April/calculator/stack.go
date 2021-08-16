package main

import "errors"

//Stack's float implementation.
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

//Stack's string implementation.
type StrStack []string

func (str *StrStack) Push(value string) {
	*str = append(*str, value)
}

func (str *StrStack) Pop() (string, error) {
	temp, err := str.Peek()
	if err != nil {
		return "", err
	}

	*str = (*str)[0 : len(*str)-1]

	return temp, nil
}

func (str *StrStack) Peek() (string, error) {
	if len(*str) == 0 {
		return "", errors.New("stack is empty")
	}

	return (*str)[len(*str)-1], nil
}

func (str *StrStack) Len() int {
	return len(*str)
}
