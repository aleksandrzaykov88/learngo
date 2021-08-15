package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type function struct {
	numParams int                          // the number of arguments needed from stack
	result    func(args []float64) float64 // the function to execute to obtain the result
}

// regular expression to ensure it's an int literal
var numberExp = regexp.MustCompile(`^[+-]?\d+(\.\d+)?$`) // matches "2.75", "+129832", "-843", etc.

// map of functions for operators
//
// example explanation for the + function:
// numParams = 2 means eval() will pass it the top 2 values from the stack
// the result function simply casts the first two values from the slice of values it got passed to int, adds them up, and returns the result
//
// here is a formatted version of the + line:
//	"+": function{
//		numParams: 2,
//		result: func(args []float64) float64 {
//			return args[1] + args[0]
//		},
//	},
var functions = map[string]function{
	"+": function{2, func(args []float64) float64 { return args[1] + args[0] }},                        // addition
	"·": function{2, func(args []float64) float64 { return args[1] * args[0] }},                        // multiplication
	"-": function{2, func(args []float64) float64 { return args[1] - args[0] }},                        // subtraction
	"/": function{2, func(args []float64) float64 { return args[1] / args[0] }},                        // division
	"%": function{2, func(args []float64) float64 { return float64(int64(args[1]) % int64(args[0])) }}, // modulo
	"^": function{2, func(args []float64) float64 { return math.Pow(args[1], args[0]) }},               // squaring
}

func eval(token string) {
	// the arguments we will later pass to the function
	args := []float64{}

	// choose correct function
	f, ok := functions[token]
	if !ok {
		// probably dealing with an int literal, but let's make sure
		if !numberExp.MatchString(token) {
			// invalid input
			fmt.Fprintln(os.Stderr, "invalid input:", token)
			os.Exit(1)
		}

		f = function{0, func(args []float64) float64 { return args[0] }} // for int literals; returns its only argument

		// because x is an int literal, convert to int and add it as argument
		num, err := strconv.ParseFloat(token, 64)
		if err != nil {
			panic(err)
		}

		args = append(args, num)
	}

	// pop needed arguments from stack
	for i := 0; i < f.numParams; i++ {
		x, err := s.Pop()
		if err != nil {
			panic(err)
		}

		args = append(args, x)
	}

	// push result back onto stack
	s.Push(f.result(args))
}
