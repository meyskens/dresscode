package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var buildOrder = []string{ // or else it will end weird...
	"Body",
	"Eyes",
	"Shirts",
	"Hair",
	"Glasses",
	"Hats_and_Hair_Accessories",
	"Extras",
}

const input = ""

var guard = make(chan struct{}, 5)

func main() {
	words := strings.Split(input, " ")
	files := []string{}
	donechan := make(chan bool)
	toBuild := unique(words)

	for i, word := range toBuild {
		guard <- struct{}{}
		go createGopher(i, word, donechan)
		files = append(files, fmt.Sprintf("%d.png", i))
		fmt.Printf("Started %d/%d\n", i+1, len(toBuild))
	}

	for i := 0; i < len(toBuild); i++ {
		<-donechan
	}
	fmt.Println("Making collage")
	collage(files)
}

func createGopher(i int, word string, out chan bool) {
	defer func() { out <- true }()
	url, err := createForWord(word)
	if err != nil {
		return
	}
	f, err := os.Create(fmt.Sprintf("%d.png", i))
	if err != nil {
		return
	}
	defer f.Close()
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	io.Copy(f, resp.Body)

	<-guard
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
