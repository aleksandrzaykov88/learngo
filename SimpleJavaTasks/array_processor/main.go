package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//Arr describes the simple array.
type Arr []int

//printArr prints array.
func (a *Arr) printArr() {
	fmt.Println(*a)
}

//inputArr gets array from user.
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

//sumArr returns the sum of all array elements.
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

//evenAmount returns amount of even numbers among array elements.
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

//abSlice returns amount of numbers in array which are between "from" and "to" nums.
func (a *Arr) abSlice(from, to int) int {
	count := 0
	for _, num := range *a {
		if num >= from && num <= to {
			count++
		}
	}
	return count
}

//positiveCheck returns true if all elements in array are positive.
func (a *Arr) positiveCheck() bool {
	for _, num := range *a {
		if num <= 0 {
			return false
		}
	}
	return true
}

//reverseArr reverses the order of elements in array.
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
