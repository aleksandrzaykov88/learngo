package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"runtime"
)

const result = "Polar radius=%.02f q=%.02f degrees (r) Cartesian x=%.02f y=%.02f\n"

type polar struct {
	radius float64
	Teta   float64
}

type cartesian struct {
	x float64
	y float64
}

var prompt = "Enter a raddius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

//init determines the contents of the prompt string, considering differences in end-of-file designation on different platforms.
func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

//createSolver creates an answers channel and after that sends responses through it.
func createSolver(questions chan polar) chan cartesian {
	answers := make(chan cartesian)
	go func() {
		for {
			polarCoord := <-questions
			Teta := polarCoord.Teta * math.Pi / 180.0
			x := polarCoord.radius * math.Cos(Teta)
			y := polarCoord.radius * math.Sin(Teta)
			answers <- cartesian{x, y}
		}
	}()
	return answers
}

//interact prompts the user to enter polar coordinates (radius and angle).
func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var radius, Teta float64
		if _, err := fmt.Sscanf(line, "%f %f", &radius, &Teta); err != nil {
			fmt.Fprintln(os.Stderr, "invalid input")
			continue
		}
		questions <- polar{radius, Teta}
		coord := <-answers
		fmt.Printf(result, radius, Teta, coord.x, coord.y)
	}
	fmt.Println()
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}
