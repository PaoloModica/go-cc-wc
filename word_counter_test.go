package wordcounter_test

import (
	"fmt"
	wordcounter "go-cc-wc"
	"reflect"
	"testing"
)

func TestWC(t *testing.T) {
	var wordCounter = wordcounter.WordCounterFunc(wordcounter.WordCount)

	type wcTestCase struct {
		id             string
		option         string
		expectedResult []int
	}

	t.Run("raise an error for invalid option", func(t *testing.T) {
		file := "test.txt"
		option := "-y"

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
		wcTestCases := []wcTestCase{
			{
				"-c option returns number of bytes in file",
				"-c",
				[]int{335043},
			},
			{
				"-l option returns number of lines in file",
				"-l",
				[]int{7146},
			},
			{
				"-w option returns number of words in file",
				"-w",
				[]int{58164},
			},
			{
				"-m option returns number of characters in file",
				"-m",
				[]int{332147},
			},
			{
				"no option returns the number provided by -l, -w and -c option",
				"",
				[]int{7146, 58164, 335043},
			},
		}

		file := "test.txt"

		for _, testCase := range wcTestCases {
			t.Run(fmt.Sprintf("wc with %s", testCase.id), func(t *testing.T) {
				gotResult, err := wordCounter.CountWords(file, testCase.option)

				if err != nil {
					t.Fatalf("could not read the content of %s, error: %v", file, err)
				}

				if !reflect.DeepEqual(gotResult, testCase.expectedResult) {
					t.Errorf("expected %v, got %v with option %s", testCase.expectedResult, gotResult, testCase.option)
				}

			})
		}
	})
}
