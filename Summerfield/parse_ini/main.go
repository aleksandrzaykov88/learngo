package main

import "fmt"

//ParseIni parses .ini-file.
func ParseIni() map[string]map[string]string {
	return map[string]map[string]string{"Map": {"Key": "Value"}}
}

func main() {
	fmt.Println(ParseIni())
}
