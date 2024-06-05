package main

import "testing"

func TestCalculateByteCount(t *testing.T) {

	got := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")

	expected := 342190

	if got != expected {
		t.Errorf("got %d but expected %d", got, expected)
	}

}
