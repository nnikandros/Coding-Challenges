package main

import (
	"reflect"
	"testing"
)

func TestHandleArguments(t *testing.T) {

	t.Run("Run with only -c flag", func(t *testing.T) {
		exampleArgs := []string{"-c", "test.txt"}
		got := HandleArguments(exampleArgs)

		expected := ParsedArgs{flagOptions: []string{"-c"}, name: "test.txt"}
		// assertCorrectResult(t, got, expected)
		assertCorrectStruct(t, got, expected)
	})

	t.Run("Run with only -l flag", func(t *testing.T) {
		exampleArgs := []string{"-l", "test.txt"}
		got := HandleArguments(exampleArgs)

		expected := ParsedArgs{flagOptions: []string{"-l"}, name: "test.txt"}
		assertCorrectStruct(t, got, expected)
	})

	t.Run("Run with no flags. Default mode is -c -l -w", func(t *testing.T) {
		exampleArgs := []string{"test.txt"}
		got := HandleArguments(exampleArgs)
		expected := ParsedArgs{flagOptions: []string{"-c", "-l", "-w"}, name: "test.txt"}
		assertCorrectStruct(t, got, expected)
	})

}

func assertCorrectStruct(t testing.TB, got, expected ParsedArgs) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v but expected %v", got, expected)
	}

}
