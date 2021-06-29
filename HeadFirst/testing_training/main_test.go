package main

import (
	"fmt"
	"testing"
)

func errorString(list []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", want \"%s\"", list, got, want)
}

func TestTwoElements(t *testing.T) {
	list := []string{"apple", "orange"}
	want := "apple and orange"
	got := joinWithCommas(list)
	if want != got {
		t.Error(errorString(list, got, want))
	}
}

func TestThreeElements(t *testing.T) {
	list := []string{"apple", "orange", "pear"}
	want := "apple, orange, and pear"
	got := joinWithCommas(list)
	if got != want {
		t.Error(errorString(list, got, want))
	}
}
