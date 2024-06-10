package main

import "testing"

func TestCountChars(t *testing.T) {

	got := CountChars("C:\\Users\\nnika\\Desktop\\projects\\coding-challenges\\build-wc\\test.txt")

	expected := 339292

	assertCorrectResult(t, got, expected)

}
