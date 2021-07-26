package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//toDigit converts string go digit.
func toDigit(s string) int64 {
	r := []rune(s)
	res := ""
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(r[i]) {
			res += string(r[i])
		}
	}
	num, _ := strconv.Atoi(string(res))
	return int64(num)
}

//splitNum splits string by ','-sing on whole and fractional numbers.
func splitNum(num string) (w, f string) {
	s := strings.Split(num, ",")
	if len(s) == 1 {
		return s[0], "0"
	}
	w, f = s[0], s[1]
	return
}

//getF creates float number from string whitch are its whole and fractional parts.
func getF(num string) float64 {
	w, f := splitNum(num)
	wn := float64(toDigit(w)) + float64(toDigit(f))/math.Pow(10, float64(len(f)))
	return wn
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	nums := strings.Split(str, ";")
	n1, n2 := nums[0], nums[1]
	fmt.Printf("%.4f", getF(n1)/getF(n2))
}
