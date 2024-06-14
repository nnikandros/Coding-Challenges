package main

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestHandleArguments(t *testing.T) {

	t.Run("Run with only -c flag", func(t *testing.T) {
		exampleArgs := []string{"-c", pathToTestFile()}
		got := HandleArguments(exampleArgs)

		fileInfo, _ := os.Open(pathToTestFile())

		expected := ParsedArgs{flagOptions: []string{"-c"}, path: pathToTestFile(), fileName: fileInfo.Name()}
		// assertCorrectResult(t, got, expected)
		assertCorrectStruct(t, got, expected)
	})

	t.Run("Run with only -l flag", func(t *testing.T) {
		exampleArgs := []string{"-l", pathToTestFile()}
		got := HandleArguments(exampleArgs)

		fileInfo, _ := os.Open(pathToTestFile())
		expected := ParsedArgs{flagOptions: []string{"-l"}, path: pathToTestFile(), fileName: fileInfo.Name()}
		assertCorrectStruct(t, got, expected)
	})

	t.Run("Run with no flags. Default mode is -c -l -w", func(t *testing.T) {
		exampleArgs := []string{pathToTestFile()}
		got := HandleArguments(exampleArgs)

		fileInfo, _ := os.Open(pathToTestFile())

		expected := ParsedArgs{flagOptions: []string{"-c", "-l", "-w"}, path: pathToTestFile(), fileName: fileInfo.Name()}
		assertCorrectStruct(t, got, expected)
	})

}

func TestDebugging(t *testing.T) {

	t.Run("debugging test", func(t *testing.T) {

		fileInfo, _ := os.Open(pathToTestFile())

		fmt.Println(fileInfo)

	})

}

func assertCorrectStruct(t testing.TB, got, expected ParsedArgs) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %v but expected %v", got, expected)
	}

}
