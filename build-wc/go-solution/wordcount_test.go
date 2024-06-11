package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got := CountWords(pathToTestFile())
		// fmt.Println(got)
		expected := 58164
		assertCorrectResult(t, got, expected)

	})
}
