package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCalculateByteCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got, _ := CalculateByteCount(pathToTestFile())

		expected := 342190
		assertCorrectResult(t, got, expected)

	})

	// t.Run("relative path", func(t *testing.T) {
	// 	got, _ := CalculateByteCount("../build-wc/test.txt")

	// 	expected := 342190
	// 	assertCorrectResult(t, got, expected)
	// })
}

func assertCorrectResult(t testing.TB, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}
}

func pathToTestFile() string {
	working_dir, _ := os.Getwd()
	// path uses Unix for paths and does not deal with windows paths. use filepath dir
	return filepath.Join(filepath.Dir(working_dir), "test.txt")

}
