package main

import "testing"

func TestHandleArguments(t *testing.T) {

	t.Run("run with arguments", func(t *testing.T) {
		exampleArgs := []string{"-c", "test.txt"}
		HandleArguments(exampleArgs)

	})

}
