package main

import "fmt"

//FuzzyBool is a custom type for display bool values in float-view(like 0.25 == false, 0.75 == true).
type FuzzyBool struct {
	value float32
}

//New returns the new object of FuzzyBool type.
func New(value interface{}) (*FuzzyBool, error) {
	amount, err := float32ForValue(value)
	return &FuzzyBool{amount}, err
}

//float32ForValue converts input in float32 number or returns the error.
func float32ForValue(value interface{}) (fuzzy float32, err error) {
	switch value := value.(type) {
	case float32:
		fuzzy = value
	case float64:
		fuzzy = float32(value)
	case int:
		fuzzy = float32(value)
	case bool:
		fuzzy = 0
		if value {
			fuzzy = 1
		}
	default:
		return 0, fmt.Errorf("float32ForValue(): %v is not a "+"number or Boolean", value)
	}
	if fuzzy < 0 {
		fuzzy = 0
	} else if fuzzy > 1 {
		fuzzy = 1
	}
	return fuzzy, nil
}

//String is a stringer for FuzzyBool
func (fuzzy *FuzzyBool) String() string {
	return fmt.Sprintf("%.0f%%", 100*fuzzy.value)
}

//Set is a setter for FuzzyBool.
func (fuzzy *FuzzyBool) Set(value interface{}) (err error) {
	fuzzy.value, err = float32ForValue(value)
	return err
}

//Copy returns a pointer of new FuzzyBool.
func (fuzzy *FuzzyBool) Copy() *FuzzyBool {
	return &FuzzyBool{fuzzy.value}
}

//Not reverts value of FuzzyBool.
func (fuzzy *FuzzyBool) Not() *FuzzyBool {
	return &FuzzyBool{1 - fuzzy.value}
}

//And returns min from input values.
func (fuzzy *FuzzyBool) And(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	minimum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if minimum > other.value {
			minimum = other.value
		}
	}
	return &FuzzyBool{minimum}
}

//Or returns max from input values.
func (fuzzy *FuzzyBool) Or(first *FuzzyBool, rest ...*FuzzyBool) *FuzzyBool {
	maximum := fuzzy.value
	rest = append(rest, first)
	for _, other := range rest {
		if maximum < other.value {
			maximum = other.value
		}
	}
	return &FuzzyBool{maximum}
}

//Less returns true if input value > *FuzzyBool value.
func (fuzzy *FuzzyBool) Less(other *FuzzyBool) bool {
	return fuzzy.value < other.value
}

//Equal returns true if input value == *FuzzyBool value.
func (fuzzy *FuzzyBool) Equal(other *FuzzyBool) bool {
	return fuzzy.value == other.value
}

//Bool is a function for type casting into bool value.
func (fuzzy *FuzzyBool) Bool() bool {
	return fuzzy.value >= .5
}

//Float is a function for type casting into float value.
func (fuzzy *FuzzyBool) Float() float64 {
	return float64(fuzzy.value)
}

func main() {
	a, _ := New(0)
	b, _ := New(.25)
	c, _ := New(.75)
	d := c.Copy()
	if err := d.Set(1); err != nil {
		fmt.Println(err)
	}
	process(a, b, c, d)
	s := []*FuzzyBool{a, b, c, d}
	fmt.Println(s)
}

func process(a, b, c, d *FuzzyBool) {
	fmt.Println("Original:", a, b, c, d)
	fmt.Println("Not: ", a.Not(), b.Not(), c.Not(), d.Not())
	fmt.Println("Not Not: ", a.Not().Not(), b.Not().Not(), c.Not().Not(),
		d.Not().Not())
	fmt.Print("0.And(.25)->", a.And(b), "• .25.And(.75)->", b.And(c),
		"• .75.And(1)->", c.And(d), " • .25.And(.75,1)->", b.And(c, d), "\n")
	fmt.Print("0.Or(.25)->", a.Or(b), "• .25.Or(.75)->", b.Or(c),
		"• .75.Or(1)->", c.Or(d), " • .25.Or(.75,1)->", b.Or(c, d), "\n")
	fmt.Println("a < c, a == c, a > c:", a.Less(c), a.Equal(c), c.Less(a))
	fmt.Println("Bool: ", a.Bool(), b.Bool(), c.Bool(), d.Bool())
	fmt.Println("Float: ", a.Float(), b.Float(), c.Float(), d.Float())
}
