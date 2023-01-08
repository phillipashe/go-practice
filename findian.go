package main

import (
	"fmt"
	"strings"
)

func main() {
	input := new(string)

	/*
	 The program should print “Found!” if the entered string starts
	 with the character ‘i’, ends with the character ‘n’, and contains the character ‘a’.

	 Otherwise print "Not Found!".
	*/
	fmt.Println("Type something: ")
	fmt.Scan(input)

	startsWithI := strings.HasPrefix(*input, "i")
	hasA := strings.Contains(*input, "a")
	endsWithN := strings.HasSuffix(*input, "n")

	if startsWithI && hasA && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
