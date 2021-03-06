package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Song struct {
	Title    string
	Filename string
	Seconds  int
}

//readPlsPlaylist reads information from string received from file in .pls extension.
func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "[playlist]") || strings.HasPrefix(line, "Version") || strings.HasPrefix(line, "Number") {
			continue
		}
		separatorIndex := strings.Index(line, "=")
		if strings.HasPrefix(line, "Title") {
			song.Title = line[separatorIndex+1:]
		}
		if strings.HasPrefix(line, "Length") {
			latency, err := strconv.ParseInt(line[separatorIndex+1:], 10, 64)
			if err != nil {
				log.Fatal(err)
				break
			}
			song.Seconds = int(latency)
		}
		if strings.HasPrefix(line, "File") {
			song.Filename = line[separatorIndex+1:]
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

//writeM3uPlaylist writes the processed string in m3u-format in console.
func writeM3uPlaylist(songs []Song) {
	fmt.Println("#EXTM3U")
	for _, song := range songs {
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
		fmt.Printf("%s\n", song.Filename)
	}
}

//readM3uPlaylist reads information from string received from file in .m3u extension.
func readM3uPlaylist(data string) (songs []Song) {
	var song Song
	for _, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#EXTM3U") {
			continue
		}
		if strings.HasPrefix(line, "#EXTINF:") {
			song.Title, song.Seconds = parseExtinfLine(line)
		} else {
			song.Filename = strings.Map(mapPlatformDirSeparator, line)
		}
		if song.Filename != "" && song.Title != "" && song.Seconds != 0 {
			songs = append(songs, song)
			song = Song{}
		}
	}
	return songs
}

//writePlsPlaylist writes the processed string in pls-format in console.
func writePlsPlaylist(songs []Song) {
	fmt.Println("[playlist]")
	for i, song := range songs {
		i++
		fmt.Printf("File%d=%s\n", i, song.Filename)
		fmt.Printf("Title%d=%s\n", i, song.Title)
		fmt.Printf("Length%d=%d\n", i, song.Seconds)
	}
	fmt.Printf("NumberOfEntries=%d\nVersion=2\n", len(songs))
}

//parseExtinfLine parses m3u file to get filename, title and duration.
func parseExtinfLine(line string) (title string, seconds int) {
	if i := strings.IndexAny(line, "-0123456789"); i > -1 {
		const separator = ","
		line = line[i:]
		if j := strings.Index(line, separator); j > -1 {
			title = line[j+len(separator):]
			var err error
			if seconds, err = strconv.Atoi(line[:j]); err != nil {
				log.Printf("failed to read the duration for '%s': %v\n", title, err)
				seconds = -1
			}
		}
	}
	return title, seconds
}

//mapPlatformDirSeparator returns separator sign for the curren OS.
func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

func main() {
	var m3uF, plsF bool
	if len(os.Args) == 1 || !strings.HasSuffix(os.Args[1], ".m3u") && !strings.HasSuffix(os.Args[1], ".pls") {
		fmt.Printf("usage: %s <file.m3u> or <file.pls>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if strings.HasSuffix(os.Args[1], ".pls") {
		plsF = true
	} else {
		m3uF = true
	}
	if rawBytes, err := ioutil.ReadFile(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		if m3uF {
			songs := readM3uPlaylist(string(rawBytes))
			writePlsPlaylist(songs)
		} else if plsF {
			songs := readPlsPlaylist(string(rawBytes))
			writeM3uPlaylist(songs)
		}
	}
}
