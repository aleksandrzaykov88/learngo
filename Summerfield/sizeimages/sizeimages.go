package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// htmlImageFormat reads html-file and scans it for img-tags.
// If in it is no width/height image data
// htmlImageFormat calls getImageData func.
func htmlImageFormat(filename string) {
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
				getImageData(matchAttr[1])
				// If in img-tag no width/height attributes
				// Write them right in html-file.
				resW := strings.Contains(matchTag[0], "width")
				if !resW {
					lines[i] = "LOL"
				}
				resH := strings.Contains(matchTag[0], "height")
				if !resH {
					lines[i] = "LOL HEIGHT"
				}
			}
		}
	}

	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}

func getImageData(link string) {
	// Get the data
	resp, err := http.Get(link)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	config, format, err := image.DecodeConfig(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(link, format, config.Height, config.Width, config.ColorModel)
	/*
		// Create the file
		out, err := os.Create(`C:\Users\aleks\Documents\GitHub\learngo\Summerfield\sizeimages\download.jpeg`)
		if err != nil {
			log.Panic(err)
		}
		defer out.Close()

		// Write the body to file
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.Open(out.Name())
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		config, format, err := image.DecodeConfig(bufio.NewReader(file))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(format, config.Height, config.Width, config.ColorModel)*/
}

func main() {
	htmlImageFormat("test.html")
}
