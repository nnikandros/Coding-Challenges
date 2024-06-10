package main

import (
	"fmt"
	"testing"
)

func TestWordCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got := CountWords("hello world simple        malakas\nkapikos")
		fmt.Println(got)
		// expected := 342190
		// assertCorrectResult(t, got, expected)

	})
}
