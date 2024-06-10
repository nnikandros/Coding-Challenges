package main

import (
	"testing"
)

func TestWordCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got := CountWords("C:\\Users\\nnika\\Desktop\\projects\\coding-challenges\\build-wc\\test.txt")
		// fmt.Println(got)
		expected := 58164
		assertCorrectResult(t, got, expected)

	})
}
