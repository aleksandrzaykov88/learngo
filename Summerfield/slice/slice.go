package main

import "log"

type Slice struct {
	slice []interface{}
	less  func(interface{}, interface{}) bool
}

func New(less func(interface{}, interface{}) bool) *Slice {
	return &Slice{less: less}
}

func NewStringSlice() *Slice {
	return &Slice{less: func(a, b interface{}) bool {
		return a.(string) < b.(string)
	}}
}

func NewIntSlice() *Slice {
	return &Slice{less: func(a, b interface{}) bool {
		return a.(int) < b.(int)
	}}
}

func (s *Slice) Clear() {
	s.slice = s.slice[:0]
}

func (s *Slice) Add(x interface{}) {
	index := bisectLeft(s.slice, s.less, x)
	if index == len(s.slice) || len(s.slice) == 0 {
		s.slice = append(s.slice, x)
		return
	}
	newSlice := make([]interface{}, 0)
	newSlice = append(newSlice, s.slice[:index]...)
	newSlice = append(newSlice, x)
	newSlice = append(newSlice, s.slice[index:]...)
	s.slice = newSlice
}

func (s *Slice) Remove(input interface{}) bool {
	index := s.Index(input)
	if index == -1 {
		return false
	}
	newSlice := make([]interface{}, 0)
	newSlice = append(newSlice, s.slice[:index]...)
	newSlice = append(newSlice, s.slice[index+1:]...)
	s.slice = newSlice
	return true
}

func (s *Slice) Index(input interface{}) int {
	for i, val := range s.slice {
		if val == input {
			return i
		}
	}
	return -1
}

func (s *Slice) At(index int) interface{} {
	if index >= len(s.slice) {
		log.Panic("Index in out of range.")
	}
	return s.slice[index]
}

func (s *Slice) Len() int {
	return len(s.slice)
}

func bisectLeft(slice []interface{}, less func(interface{}, interface{}) bool, x interface{}) int {
	left, right := 0, len(slice)
	for left < right {
		middle := int((left + right) / 2)
		if less(slice[middle], x) {
			left = middle + 1
		} else {
			right = middle
		}
	}
	return left
}
