package main

import (
	"fmt"
	"log"
)

type Font struct {
	family string
	size   int
}

func CheckSize(size int) bool {
	if size < 5 || size > 144 {
		log.Printf("%d is invalid\n", size)
		return false
	}
	return true
}

func CheckFamily(family string) bool {
	if family == "" {
		log.Printf("%s is invalid\n", family)
		return false
	}
	return true
}

func New(family string, size int) *Font {
	var f = "Arial"
	var s = 12
	if CheckSize(size) {
		s = size
	}
	if CheckFamily(family) {
		f = family
	}
	return &Font{f, s}
}

func (f *Font) SetFamily(family string) {
	if !CheckFamily(family) {
		return
	}
	f.family = family
}

func (f *Font) SetSize(size int) {
	if !CheckSize(size) {
		return
	}
	f.size = size
}

func (f *Font) Family() string {
	return f.family
}

func (f *Font) Size() int {
	return f.size
}

func (f *Font) String() string {
	return fmt.Sprintf(`{font-family: "%v"; font-size: %vpt;}`, f.family, f.size)
}
