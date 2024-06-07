package main

import "testing"

func TestCalculateByteCount(t *testing.T) {
	t.Run("absolute path", func(t *testing.T) {
		got, _ := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")

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
