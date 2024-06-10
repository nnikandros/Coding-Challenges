package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
)

func CalculateLineCount(name string) int {
	file, _ := os.Open(name)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	return lineCount
	// count, err := LineCounter(file)
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }

	// return count
}

func LineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
