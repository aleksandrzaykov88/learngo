package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"
)

const maxGoroutines = 100
const maxSizeOfSmallFile = 1024 * 32

type fileInfo struct {
	sha1 []byte
	size int64
	path string
}

type pathsInfo struct {
	size  int64
	paths []string
}

//findDuplicates findes duplicates in dirname
func findDuplicates(infoChan chan fileInfo, dirname string) {
	waiter := &sync.WaitGroup{}
	filepath.Walk(dirname, makeWalkFunc(infoChan, waiter))
	waiter.Wait()
	close(infoChan)
}

//makeWalkFunc returns a WalkFunc
func makeWalkFunc(infoChan chan fileInfo, waiter *sync.WaitGroup) func(string, os.FileInfo, error) error {
	return func(path string, info os.FileInfo, err error) error {
		if err == nil && info.Size() > 0 && (info.Mode()&os.ModeType == 0) {
			if info.Size() < maxSizeOfSmallFile || runtime.NumGoroutine() > maxGoroutines {
				processFile(path, info, infoChan, nil)
			} else {
				waiter.Add(1)
				go processFile(path, info, infoChan, func() { waiter.Done() })
			}
		}
		return nil
	}
}

//processFile calculates SHA-1 hash sum
func processFile(filename string, info os.FileInfo, infoChan chan fileInfo, done func()) {
	if done != nil {
		defer done()
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Println("error:", err)
		return
	}
	defer file.Close()
	hash := sha1.New()
	if size, err := io.Copy(hash, file); size != info.Size() || err != nil {
		if err != nil {
			log.Println("error:", err)
		} else {
			log.Println("error: failed to read the whole file:", filename)
		}
		return
	}
	infoChan <- fileInfo{hash.Sum(nil), info.Size(), filename}
}

//mergeResults creates and returns result-map
func mergeResults(infoChan <-chan fileInfo) map[string]*pathsInfo {
	pathData := make(map[string]*pathsInfo)
	format := fmt.Sprintf("%%016X:%%%dX", sha1.Size*2)
	for info := range infoChan {
		key := fmt.Sprintf(format, info.size, info.sha1)
		value, found := pathData[key]
		if !found {
			value = &pathsInfo{size: info.size}
			pathData[key] = value
		}
		value.paths = append(value.paths, info.path)
	}
	return pathData
}

//outputResults outputs duplicates in path
func outputResults(pathData map[string]*pathsInfo) {
	keys := make([]string, 0, len(pathData))
	for key := range pathData {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		value := pathData[key]
		if len(value.paths) > 1 {
			fmt.Printf("%d duplicate files (%s bytes):\n", len(value.paths), commas(value.size))
			sort.Strings(value.paths)
			for _, name := range value.paths {
				fmt.Printf("\t%s\n", name)
			}
		}
	}
}

//commas adds commas to output
func commas(x int64) string {
	value := fmt.Sprint(x)
	for i := len(value) - 3; i > 0; i -= 3 {
		value = value[:i] + "," + value[i:]
	}
	return value
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage: %s <path>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	infoChan := make(chan fileInfo, maxGoroutines*2)
	go findDuplicates(infoChan, os.Args[1])
	pathData := mergeResults(infoChan)
	outputResults(pathData)
}
