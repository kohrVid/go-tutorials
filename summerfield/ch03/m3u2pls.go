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

func main() {
	if len(os.Args) == 1 {
		usageMessage()
	} else if strings.HasSuffix(os.Args[1], ".m3u") {
		fromM3u(os.Args[1])
	} else if strings.HasSuffix(os.Args[1], ".pls") {
		fromPls(os.Args[1])
	} else {
		usageMessage()
	}

}

func usageMessage() {
	fmt.Printf("usage: %s <file.[pls|m3u]>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func fromM3u(file_name string) {
	if rawBytes, err := ioutil.ReadFile(file_name); err != nil {
		log.Fatal(err)
	} else {
		songs := readM3uPlaylist(string(rawBytes))
		writePlsPlaylist(songs)
	}
}

func fromPls(file_name string) {
	if rawBytes, err := ioutil.ReadFile(file_name); err != nil {
		log.Fatal(err)
	} else {
		songs := readPlsPlaylist(string(rawBytes))
		fmt.Println(songs)
		writeM3uPlaylist(songs)
	}
}

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

func readPlsPlaylist(data string) (songs []Song) {
	var song Song
	for i, line := range strings.Split(data, "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Title") {
			song.Title = line[7+len(string(i)):]
		} else if strings.HasPrefix(line, "File") {
			song.Filename = line[6+len(string(i)):]
		} else if strings.HasPrefix(line, "Length") {
			song.Seconds, _ = strconv.Atoi(line[8+len(string(i)):])
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

func mapPlatformDirSeparator(char rune) rune {
	if char == '/' || char == '\\' {
		return filepath.Separator
	}
	return char
}

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

func writeM3uPlaylist(songs []Song) {
	for i, song := range songs {
		i++
		fmt.Println(song.Filename)
		fmt.Printf("#EXTINF:%d,%s\n", song.Seconds, song.Title)
	}
}
