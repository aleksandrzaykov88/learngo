package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var bigDigits = [][]string{
	{"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  "},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ",
		"   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}

//bigDigitsPrinter get number from command line and after prints it lik BIG DIGIT.
func bigDigitsPrinter() {
	var barFlag, helpFlag bool
	for _, arg := range os.Args {
		if arg == "-b" || arg == "--bar" {
			barFlag = true
		} else if arg == "-h" || arg == "--help" {
			helpFlag = true
		}
	}
	if len(os.Args) == 1 || helpFlag {
		fmt.Printf("usage: %s <whole-number>\n-b --bar draw an underbar and an overbar\n", filepath.Base(os.Args[0]))
		if len(os.Args) == 1 {
			os.Exit(1)
		}
	}
	stringOfDigits := os.Args[1]
	starSum := -1

	for num := range stringOfDigits {
		digit := stringOfDigits[num] - '0'
		if 0 <= digit && digit <= 9 {
			starSum += len(bigDigits[digit][0])
		} else {
			log.Fatal("invalid whole number")
		}
	}
	starSum += len(stringOfDigits)
	if barFlag {
		for i := 0; i < starSum; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			}
		}
		fmt.Println(line)
	}

	if barFlag {
		for i := 0; i < starSum; i++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	bigDigitsPrinter()
}
