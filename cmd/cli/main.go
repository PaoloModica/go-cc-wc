package main

import (
	"fmt"
	wordcounter "go-cc-wc"
	"os"
)

func main() {
	var file string
	var option string

	if len(os.Args) > 2 {
		option = os.Args[1]
		file = os.Args[2]
	} else {
		file = os.Args[1]
	}
	wordCounter := wordcounter.WordCounterFunc(wordcounter.WordCount)

	// feed the word counter with the request args
	res, _ := wordCounter.CountWords(file, option)
	fmt.Printf("%+v %s\n", res, file)
}
