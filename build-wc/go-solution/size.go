package main

import (
	"fmt"
	"os"
)

func CalculateByteCount(name string) int {

	bytes, err := os.ReadFile(name)
	if err != nil {
		z := fmt.Errorf("error happend with input %s", name)

		fmt.Println(z.Error())

		return 0, fmt.Errorf("error reading file %s: %v", name, err)
	}

	return len(bytes), nil

}
