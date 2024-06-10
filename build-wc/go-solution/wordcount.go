package main

import "strings"

func CountWords(name string) []string {
	return strings.Split(name, " ")
}
