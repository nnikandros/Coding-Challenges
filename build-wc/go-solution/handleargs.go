package main

import (
	"fmt"
)

type ParsedArgs struct {
	flagOptions []string
	path        string
	fileName    string
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

	// here return an error if the  file doesnt exis

	if len(cli_args) == 1 {

		return ParsedArgs{flagOptions: []string{"-c", "-l", "-w"}, path: last}
	}

	return ParsedArgs{flagOptions: uptoLast, path: last}
}
