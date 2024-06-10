package main

import (
	"testing"
)

func TestCalculateLineCount(t *testing.T) {
	t.Run("line count with absolute path", func(t *testing.T) {
		// got := CalculateLineCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")

		got := CalculateLineCount("C:\\Users\\nnika\\Desktop\\projects\\coding-challenges\\build-wc\\test.txt")

		// fmt.Println(got)

		expected := 7145
		assertCorrectResult(t, got, expected)
	})

}
