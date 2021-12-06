package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const (
	w = "width"
	h = "height"
)

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	argFiles := os.Args[1:] // Get filenames from command args
	if len(argFiles) > 0 {
		for _, filename := range argFiles {
			lines := make(chan string, workers*4) // Creating workers pool
			done := make(chan struct{}, workers)

			go htmlImageFormat(filename, lines)
			processLines(done, lines)
			waitUntil(done)
			close(done)
		}
	}
}

// htmlImageFormat reads html-file
func htmlImageFormat(filename string, chLines chan<- string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		chLines <- line
	}
	close(chLines)

	// TODO:
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// processLines scans input lines for img-tags.
// If in it is no width/height image data
// processLines calls getImageData func.
func processLines(done chan<- struct{}, lines <-chan string) {
	for i := 0; i < workers; i++ {
		go func() {
			for line := range lines {
				reTag := regexp.MustCompile(`<[iI][mM][gG][^>]+>`) // Search img-tag
				matchTag := reTag.FindStringSubmatch(line)
				if len(matchTag) > 0 {
					reAttr := regexp.MustCompile(`src=["’]([^"’]+)["’]`) //Search source from img-tag
					matchAttr := reAttr.FindStringSubmatch(matchTag[0])
					if len(matchAttr) >= 1 && strings.HasPrefix(matchAttr[1], "http") {
						width, height, _ := getImageData(matchAttr[1])
						// If in img-tag no width/height attributes
						// Write them right in html-file.
						resW := strings.Contains(matchTag[0], w)
						if !resW {
							fmt.Println(addHtmlAttr(line, w, width))
						}
						resH := strings.Contains(matchTag[0], h)
						if !resH {
							fmt.Println(addHtmlAttr(line, h, height))
						}
					}
				}
			}
			done <- struct{}{}
		}()
	}
}

// getImageData checks img-link and get image properties.
// Rerurns img's width, height and error.
func getImageData(link string) (int, int, error) {
	var width, height int
	// Get the data
	resp, err := http.Get(link)
	if err != nil {
		return width, height, err
	}
	defer resp.Body.Close()

	config, _, err := image.DecodeConfig(resp.Body)
	if err != nil {
		return width, height, err
	}
	return config.Width, config.Height, err
}

// addHtmlAttr adds in img-tag width or height attribute
func addHtmlAttr(link string, attr string, value int) string {
	s := strings.Split(link, `>`)
	strValue := strconv.Itoa(value)
	fullAttr := attr + "=" + `"` + strValue + `"`
	fLink := strings.Join(s[:len(s)-1], "")
	sLink := strings.Join(s[len(s)-1:], "")

	return fLink + " " + fullAttr + ">" + sLink
}

func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}
