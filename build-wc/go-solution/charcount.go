package main

import (
	"bufio"
	"os"
)

func CountChars(name string) int {
	file, _ := os.Open(name)
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanRunes)

	// Count the words.
	count := 0
	for fileScanner.Scan() {
		count++
	}

	return count
}
