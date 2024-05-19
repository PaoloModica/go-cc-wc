package wordcounter

import (
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

func FileBytesCount(filePath string, option string) (int, error) {
	switch option {
	case "-c":
		b, err := os.ReadFile(filePath)
		if err != nil {
			return 0, errors.New(FileNotFoundErrorMessage)
		}
		return len(b), nil
	default:
		return 0, errors.New(InvalidOptionErrorMessage)
	}
}
