package main

import (
	"fmt"
	wordcounter "go-cc-wc"
	"os"
)

func main() {
	option := os.Args[1]
	file := os.Args[2]
	wordCounter := wordcounter.WordCounterFunc(wordcounter.FileBytesCount)

	// feed the word counter with the request args
	res, _ := wordCounter.CountWords(file, option)
	fmt.Printf("%d\n", res)
}
