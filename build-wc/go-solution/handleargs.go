package main

import (
	"fmt"
	"os"
)

type ParsedArgs struct {
	flagOptions []string
	name        string
}

// var possibleFlags = []string{"-c", "-l", "-m", "-w"}

// var possibleFlags2 = map[string]bool{"-c": true, "-l": true, "-w": true, "-m": true}

// Need to check the
/*
1. if the cli args is only one it needs to have the test.txt or that a file exists with this
2. if only the name of the file is given then it should return -c -l -w options
3. Checks if the file exists
*/

func HandleArguments(cli_args []string) ParsedArgs {

	//Raise an error here
	if len(cli_args) == 0 {
		fmt.Println("no name of file was provided")
	}
	last, uptoLast := cli_args[len(cli_args)-1], cli_args[:len(cli_args)-1]

	// if !strings.Contains(last, "test.txt") {

	// }
	fileInfo, err := os.Open(last)

	// here return an error if the  file doesnt exist
	if err != nil {
		fmt.Printf("Error when given the file name %s, with error %v ", last, err)
	}

	fmt.Println(fileInfo)

	fmt.Println(uptoLast)
	fmt.Println(last)

	return ParsedArgs{flagOptions: uptoLast, name: last}
}
