package main

import "testing"

func TestCalculateLineCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got := CalculateLineCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")

		expected := 7145
		assertCorrectResult(t, got, expected)
	})

}
