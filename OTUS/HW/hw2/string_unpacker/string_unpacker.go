package stringunpacker

import (
	"log"
	"strconv"
	"unicode"
)

func UnpackString(str string) string {
	if str == "" {
		return str
	}

	var (
		digit       int
		prevRune    rune
		isNumber    bool
		newString   []rune
		err         error
		isPrevSlash bool
	)
	isNumber = true // Checking if input string is not a number
	for _, r := range str {
		if !unicode.IsDigit(r) {
			isNumber = false
		}
	}
	if isNumber { // If string is number - this is incorrect string
		return ""
	}

	for i, r := range str {
		strPrev := string(prevRune)
		switch {
		case unicode.IsDigit(r) && prevRune != 0:
			digit, err = strconv.Atoi(string(r))
			if err != nil {
				log.Fatalf("Error while ATOI string %s", string(r))
			}
			if strPrev == "\\" && !isPrevSlash {
				prevRune = r
				isPrevSlash = true
			} else {
				for i := 0; i < digit; i++ {
					newString = append(newString, prevRune)
				}
				prevRune = r
				isPrevSlash = false
			}
		case prevRune != 0:
			if prevRune == r && strPrev == "\\" {
				isPrevSlash = true
			} else if strPrev != "\\" && !unicode.IsDigit(prevRune) {
				newString = append(newString, prevRune)
			} else if isPrevSlash && unicode.IsDigit(prevRune) {
				newString = append(newString, prevRune)
				isPrevSlash = false
			}
			prevRune = r
		default:
			prevRune = r
		}

		if i == len(str)-1 && !unicode.IsDigit(r) {
			newString = append(newString, r)
		} else if i == len(str)-1 && isPrevSlash && unicode.IsDigit(prevRune) {
			newString = append(newString, prevRune)
		}
	}

	return string(newString)
}
