package main

import (
	"bufio"
	"fmt"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func htmlImageFormat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		reTag := regexp.MustCompile(`<[iI][mM][gG][^>]+>`)
		matchTag := reTag.FindStringSubmatch(scanner.Text())
		if len(matchTag) > 0 {
			reAttr := regexp.MustCompile(`src=["’]([^"’]+)["’]`)
			matchAttr := reAttr.FindStringSubmatch(matchTag[0])
			if len(matchAttr) >= 1 && strings.HasPrefix(matchAttr[1], "http") {
				getImageData(matchAttr[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
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
