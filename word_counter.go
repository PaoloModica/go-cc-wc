package wordcounter

import (
	"bufio"
	"errors"
	"io"
	"os"
)

const FileNotFoundErrorMessage = "file not found"
const InvalidOptionErrorMessage = "option not valid"

type WordCounter interface {
	CountWords(filePath string, option string) (int, error)
}

type WordCounterFunc func(filePath string, option string) ([]int, error)

func (wc WordCounterFunc) CountWords(filePath string, option string) ([]int, error) {
	return wc(filePath, option)
}

func countToken(file *os.File, split bufio.SplitFunc) int {
	// set pointer to the beginning of the file
	file.Seek(0, io.SeekStart)
	scanner := bufio.NewScanner(file)
	scanner.Split(split)

	counter := 0
	for scanner.Scan() {
		counter++
	}
	return counter
}

func WordCount(filePath string, option string) ([]int, error) {
	res := []int{}
	// open the file
	file, err := os.Open(filePath)

	if err != nil {
		return res, errors.New(FileNotFoundErrorMessage)
	}

	defer file.Close()

	splits := []bufio.SplitFunc{}

	switch option {
	case "-c":
		splits = append(splits, bufio.ScanBytes)
	case "-l":
		splits = append(splits, bufio.ScanLines)
	case "-w":
		splits = append(splits, bufio.ScanWords)
	case "-m":
		splits = append(splits, bufio.ScanRunes)
	case "":
		splits = []bufio.SplitFunc{bufio.ScanLines, bufio.ScanWords, bufio.ScanBytes}
	default:
		return res, errors.New(InvalidOptionErrorMessage)
	}

	for _, split := range splits {
		res = append(res, countToken(file, split))
	}

	return res, nil
}
