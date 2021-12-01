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
	"strconv"
	"strings"
)

const (
	w = "width"
	h = "height"
)

// htmlImageFormat reads html-file and scans it for img-tags.
// If in it is no width/height image data
// htmlImageFormat calls getImageData func.
func htmlImageFormat(filename string) error {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		reTag := regexp.MustCompile(`<[iI][mM][gG][^>]+>`) // Search img-tag
		matchTag := reTag.FindStringSubmatch(line)
		if len(matchTag) > 0 {
			reAttr := regexp.MustCompile(`src=["’]([^"’]+)["’]`) //Search source from img-tag
			matchAttr := reAttr.FindStringSubmatch(matchTag[0])
			if len(matchAttr) >= 1 && strings.HasPrefix(matchAttr[1], "http") {
				width, height, err := getImageData(matchAttr[1])
				if err != nil {
					return err
				}
				// If in img-tag no width/height attributes
				// Write them right in html-file.
				resW := strings.Contains(matchTag[0], w)
				if !resW {
					lines[i] = addHtmlAttr(line, w, width)
				}
				resH := strings.Contains(matchTag[0], h)
				if !resH {
					lines[i] = addHtmlAttr(lines[i], h, height)
				}
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
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

func main() {
	err := htmlImageFormat("test.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	argFiles := os.Args[1:] // Get filenames from command args
	if len(argFiles) > 0 {
		for _, filename := range argFiles {
			htmlImageFormat(filename)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}
