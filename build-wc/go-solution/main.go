package main

import (
	"fmt"
	"os"
)

func main() {

	// file, _ := os.Open("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/go-solution/main.go")

	// Currently we assume we recieve only one argument.
	cli_args := os.Args[1:]

	// fmt.Println(cli_args)
	// possibleFlags := map[string]bool{"-c": true, "-l": true, "-w": true, "-m": true}
	// possibleFlags := []string{"-c", "-l", "-w", "-m"}
	//if no options are provided use -c -l -w
	// u := Contains(cli_args, "-c")
	// fmt.Println(u)

	// z := CalculateByteCount("/ec/local/home/nikanni/my-programming/coding-challeges/build-wc/test.txt")
	// z, err := CalculateByteCount(cli_args[1])
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%d", z)
	fmt.Println(cli_args)
}
