package wordcounter

import (
	"bufio"
	"errors"
	"os"
)

const FileNotFoundErrorMessage = "file not found"
const InvalidOptionErrorMessage = "option not valid"

type WordCounter interface {
	CountWords(filePath string, option string) (int, error)
}

type WordCounterFunc func(filePath string, option string) (int, error)

func (wc WordCounterFunc) CountWords(filePath string, option string) (int, error) {
	return wc(filePath, option)
}

func WordCount(filePath string, option string) (int, error) {
	// open the file
	file, err := os.Open(filePath)

	if err != nil {
		return 0, errors.New(FileNotFoundErrorMessage)
	}

	defer file.Close()

	var split bufio.SplitFunc

	switch option {
	case "-c":
		split = bufio.ScanBytes
	case "-l":
		split = bufio.ScanLines
	case "-w":
		split = bufio.ScanWords
	case "-m":
		split = bufio.ScanRunes
	default:
		return 0, errors.New(InvalidOptionErrorMessage)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(split)

	counter := 0
	for scanner.Scan() {
		counter++
	}

	return counter, nil
}
