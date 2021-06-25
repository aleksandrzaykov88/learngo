package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Arr []int

func (a *Arr) printArr() {
	fmt.Println(*a)
}

func (a *Arr) inputArr() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Start inputing the array.")
	fmt.Println("Press N to stop.")
	var nums []int
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		if input == "n" || input == "N" {
			break
		}
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}
	*a = nums
}

func (a *Arr) sumArr() int {
	if len(*a) == 0 {
		return 0
	}
	sum := 0
	for _, num := range *a {
		sum += num
	}
	return sum
}

func (a *Arr) evenAmount() int {
	if len(*a) == 0 {
		return 0
	}
	count := 0
	for _, num := range *a {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

func (a *Arr) abSlice(from, to int) int {
	count := 0
	for _, num := range *a {
		if num >= from && num <= to {
			count++
		}
	}
	return count
}

func (a *Arr) positiveCheck() bool {
	for _, num := range *a {
		if num <= 0 {
			return false
		}
	}
	return true
}

func (a *Arr) reverseArr() {
	var newArr, suppArr Arr
	suppArr = *a
	for i := len(suppArr) - 1; i >= 0; i-- {
		newArr = append(newArr, suppArr[i])
	}
	*a = newArr
}

func main() {
	var newArr Arr
	newArr = []int{1, 10, 11, -2}
	newArr.reverseArr()
	fmt.Println(newArr)
}
