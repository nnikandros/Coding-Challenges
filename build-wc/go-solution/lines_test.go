package main

import (
	"testing"
)

func TestCalculateLineCount(t *testing.T) {
	t.Run("line count with absolute path", func(t *testing.T) {

		got := CalculateLineCount(pathToTestFile())

		// fmt.Println(got)

		expected := 7145
		assertCorrectResult(t, got, expected)
	})

}
