package main

import (
	"fmt"
	"strconv"
)

type IPAddr [4]byte

//String is a Stringer implementation. Its prints IP with dot-separator.
func (i IPAddr) String() string {
	result := ""
	dotCounter := 0
	for _, value := range i {
		dotCounter++
		r := int(value)
		s := strconv.Itoa(r)
		if dotCounter < len(i) {
			result += s + "."
		} else {
			result += s
		}
	}
	return result
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
