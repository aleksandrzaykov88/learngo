package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//handleCommandLine handles command line args
func handleCommandLine() (algorithm int, minSize, maxSize int64, suffixes, files []string) {
	flag.IntVar(&algorithm, "algorithm", 1, "1 or 2")
	flag.Int64Var(&minSize, "min", -1, "minimum file size (-1 means no minimum)")
	flag.Int64Var(&maxSize, "max", -1, "maximum file size (-1 means no maximum)")
	var suffixesOpt *string = flag.String("suffixes", "", "comma-separated list of file suffixes")
	flag.Parse()
	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be < maximum size")
	}
	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}
	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

//source returns channel and sending to it files names
func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

//filterSuffixes it is a first filter-function in conveyor.
//Its validates filenames by their extensions
func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

//filterSize is a second filter-func.
//Its checks files for permissible size
func filterSize(minimum, maximum int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if minimum == -1 && maximum == -1 {
				out <- filename
				continue
			}
			finfo, err := os.Stat(filename)
			if err != nil {
				continue
			}
			size := finfo.Size()
			if (minimum == -1 || minimum > -1 && minimum <= size) &&
				(maximum == -1 || maximum > -1 && maximum >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

func main() {
	algorithm, minSize, maxSize, suffixes, files := handleCommandLine()
	fmt.Println(algorithm, maxSize, minSize, suffixes)
	sink(filterSize(minSize, maxSize, filterSuffixes(suffixes, source(files))))
}
