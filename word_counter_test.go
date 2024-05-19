package wordcounter_test

import (
	wordcounter "go-cc-wc"
	"testing"
)

func TestWC(t *testing.T) {
	var wordCounter = wordcounter.WordCounterFunc(wordcounter.FileBytesCount)

	t.Run("raise an error for invalid option", func(t *testing.T) {
		file := "test.txt"
		option := "-w"

		_, err := wordCounter.CountWords(file, option)

		if err == nil {
			t.Errorf("ccwc is expected to raise an error for invalid options")
		}
	})
	t.Run("raise an error if file not found", func(t *testing.T) {
		file := "unknown.txt"
		option := "-c"

		_, err := wordCounter.CountWords(file, option)

		if err == nil {
			t.Errorf("ccwc is expected to raise an error if file does not exist")
		}
	})
	t.Run("return bytes read from file", func(t *testing.T) {
		file := "test.txt"
		option := "-c"

		expectedN := 335043

		n, err := wordCounter.CountWords(file, option)

		if err != nil {
			t.Fatalf("could not read the content of %s, error: %v", file, err)
		}

		if n != expectedN {
			t.Errorf("expected %d bytes read, got %d", expectedN, n)
		}
	})
}
