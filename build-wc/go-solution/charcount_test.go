package main

import (
	"testing"
)

func TestCountChars(t *testing.T) {

	got := CountChars(pathToTestFile())

	expected := 339292

	assertCorrectResult(t, got, expected)

}
