package main

import (
	"fmt"
	"os"
)

func main() {

	// file, _ := os.Open("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/go-solution/main.go")

	arg := os.Args[1:]

	fmt.Println(arg)

	// z := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")

	// fmt.Printf("%d", z)
}
