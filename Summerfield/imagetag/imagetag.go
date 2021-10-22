package main

import (
	"fmt"
	"image"
	"os"
	"path/filepath"
	"runtime"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <image files>\n",
			filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	files := commandLineFiles(os.Args[1:])
	for _, filename := range files {
		process(filename)
	}
}

func process(filename string) {
	if info, err := os.Stat(filename); err != nil ||
		(info.Mode()&os.ModeType != 0) {
		return
	}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()
	config, _, err := image.DecodeConfig(file)
	if err != nil {
		return
	}
	fmt.Printf(`<img src="%s" width="%d" height="%d" />`,
		filepath.Base(filename), config.Width, config.Height)
	fmt.Println()
}

func commandLineFiles(files []string) []string {
	if runtime.GOOS == "windows" {
		args := make([]string, 0, len(files))
		for _, name := range files {
			if matches, err := filepath.Glob(name); err != nil {
				args = append(args, name)
			} else if matches != nil {
				args = append(args, matches...)
			}
		}
		return args
	}
	return files
}
