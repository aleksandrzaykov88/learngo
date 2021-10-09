package main

import (
	"fmt"
	"strings"
)

func getPathPrefix(charPrefix []rune, pathes [][]rune) []rune {
	if string(charPrefix) == "" {
		return []rune("")
	}
	for _, str := range pathes {
		if string(charPrefix) == string(str) || string(charPrefix) == "/" {
			return charPrefix
		}
	}
	for i := len(charPrefix) - 1; i > 0; i-- {
		if string(charPrefix[i]) != "/" {
			continue
		}
		return charPrefix[:i]
	}
	return []rune("")
}

func equalString(pathPrefix, charPrefix []rune) string {
	if string(pathPrefix) == string(charPrefix) {
		return "(==)"
	}
	return "(!=)"
}

func getEmpty(prefix []rune) []rune {
	if string(prefix) == "" {
		nprefix := string(prefix)
		nprefix = `""`
		prefix = []rune(nprefix)
	}
	return prefix
}

func main() {
	testData := [][]string{
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/goeg/prefix/extra"},
		{"/home/user/goeg", "/home/user/goeg/prefix",
			"/home/user/prefix/extra"},
		{"/pecan/π/goeg", "/pecan/π/goeg/prefix",
			"/pecan/π/prefix/extra"},
		{"/pecan/π/circle", "/pecan/π/circle/prefix",
			"/pecan/π/circle/prefix/extra"},
		{"/home/user/goeg", "/home/users/goeg",
			"/home/userspace/goeg"},
		{"/home/user/goeg", "/tmp/user", "/var/log"},
		{"/home/mark/goeg", "/home/user/goeg"},
		{"home/user/goeg", "/tmp/user", "/var/log"},
	}
	for _, strs := range testData {
		var rstrs [][]rune
		for _, str := range strs {
			rs := []rune(str)
			rstrs = append(rstrs, rs)
		}
		prefix := rstrs[0]
		len := len(prefix)
		for {
			flag := true
			for _, str := range rstrs {
				if !strings.HasPrefix(string(str), string(prefix)) {
					flag = false
					break
				}
			}
			if flag {
				break
			} else {
				len--
				prefix = prefix[:len]
			}
		}
		pathPrefix := getPathPrefix(prefix, rstrs)
		pathPrefix, prefix = getEmpty(pathPrefix), getEmpty(prefix)
		equalSign := equalString(pathPrefix, prefix)
		fmt.Println(strs)
		fmt.Println("-char:", string(prefix))
		fmt.Println(equalSign)
		fmt.Println("-path:", string(pathPrefix))
	}
}
