package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {

	// file, _ := os.Open("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/go-solution/main.go")

	arg := os.Args[1:]

	u := slices.Contains(arg, "-c")
	fmt.Println(u)
	// fmt.Println(arg)

	// z := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")
	z, err := CalculateByteCount(arg[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d", z)
}
