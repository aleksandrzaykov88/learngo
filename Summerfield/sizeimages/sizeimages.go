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

	"github.com/aleksandrzaykov88/learngo/Summerfield/sizeimages/safeslice"
)

const (
	w = "width"
	h = "height"
)

// line of file and its index
type line struct {
	index int
	text  string
}

// changeString contains old and new lines
type changeString struct {
	filename  string
	oldString string
	newString string
	index     int
}

var workers = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var ss = safeslice.New()

	argFiles := os.Args[1:] // Get filenames from command args

	fileDone := make(chan struct{})
	if len(argFiles) > 0 {
		for _, filename := range argFiles {
			go func(filename string) {
				lines := make(chan line, workers*4)
				done := make(chan struct{}, workers)

				go htmlImageFormat(filename, lines)
				processLines(done, lines, filename, ss)
				waitUntil(done)
				close(done)
				fileDone <- struct{}{}

			}(filename)
		}
		waitFileUntil(fileDone, len(argFiles))
	}

	fmt.Println(ss.Len())
}

// htmlImageFormat reads html-file and sending lines into lines channel
func htmlImageFormat(filename string, chLines chan<- line) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, l := range lines {
		chLines <- line{i, l}
	}
	close(chLines)

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

// processLines scans input lines for img-tags.
// If in it is no width/height image data
// processLines calls getImageData func.
func processLines(done chan<- struct{}, lines <-chan line, filename string, ss safeslice.SafeSlice) {
	for i := 0; i < workers; i++ {
		go func() {
			for line := range lines {
				reTag := regexp.MustCompile(`<[iI][mM][gG][^>]+>`) // Search img-tag
				matchTag := reTag.FindStringSubmatch(line.text)
				if len(matchTag) > 0 {
					reAttr := regexp.MustCompile(`src=["’]([^"’]+)["’]`) //Search source from img-tag
					matchAttr := reAttr.FindStringSubmatch(matchTag[0])
					if len(matchAttr) >= 1 && strings.HasPrefix(matchAttr[1], "http") {
						width, height, _ := getImageData(matchAttr[1])
						// If in img-tag no width/height attributes
						// Write them right in html-file.
						newString := line.text
						resW := strings.Contains(matchTag[0], w)
						if !resW {
							newString = addHtmlAttr(newString, w, width)
						}
						resH := strings.Contains(matchTag[0], h)
						if !resH {
							newString = addHtmlAttr(newString, h, height)
						}
						if !resW || !resH {
							ss.Append(changeString{filename, line.text, newString, line.index})
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

// waitUntil waits for workers
func waitUntil(done <-chan struct{}) {
	for i := 0; i < workers; i++ {
		<-done
	}
}

// waitFileUntil waits for file workers
func waitFileUntil(fileDone <-chan struct{}, filesCount int) {
	for i := 0; i < filesCount; i++ {
		<-fileDone
	}
}
