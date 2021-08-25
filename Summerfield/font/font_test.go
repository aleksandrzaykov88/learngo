package main

import (
	"fmt"
	"testing"
)

func TestFont(t *testing.T) {
	bodyFont := New("Nimbus Sans", 10)
	titleFont := New("serif", 11)
	f1(bodyFont, titleFont, t)
}

func f1(bodyFont, titleFont *Font, t *testing.T) {
	if bodyFont.String() !=
		`{font-family: "Nimbus Sans"; font-size: 10pt;}` {
		t.Fatal("#1 bodyFont invalid CSS")
	}
	if bodyFont.Size() != 10 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#2 bodyFont invalid attributes")
	}
	bodyFont.SetSize(12)
	if bodyFont.Size() != 12 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#3 bodyFont invalid attributes")
	}
	if bodyFont.String() !=
		`{font-family: "Nimbus Sans"; font-size: 12pt;}` {
		t.Fatal("#4 bodyFont invalid CSS")
	}
	bodyFont.SetFamily("")
	if bodyFont.Size() != 12 || bodyFont.Family() != "Nimbus Sans" {
		t.Fatal("#5 bodyFont invalid attributes")
	}

	if titleFont.String() != `{font-family: "serif"; font-size: 11pt;}` {
		t.Fatal("#6 titleFont invalid CSS")
	}
	if titleFont.Size() != 11 || titleFont.Family() != "serif" {
		t.Fatal("#7 titleFont invalid attributes")
	}
	titleFont.SetFamily("Helvetica")
	titleFont.SetSize(20)
	if titleFont.Size() != 20 || titleFont.Family() != "Helvetica" {
		t.Fatal("#8 titleFont invalid attributes")
	}

	f2(bodyFont, titleFont)
}

func f2(bodyFont, titleFont *Font) {
	fmt.Println(bodyFont)
	fmt.Println(titleFont)
}
