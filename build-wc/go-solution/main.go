package main

import (
	"fmt"
	"os"
)

func main() {

	// file, _ := os.Open("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/go-solution/main.go")

	// Currently we assume we recieve only one argument
	cli_args := os.Args[1:]

	u := Contains(cli_args, "-c")
	// fmt.Println(arg)

	// z := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")
	z, err := CalculateByteCount(cli_args[1])
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%d", z)
}
