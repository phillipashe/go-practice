package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

type Person struct {
	fName string
	lName string
}

func main() {
	var people []Person
	fileName := new(string)
	fmt.Println("Enter the name of the file to read from:")
	fmt.Scan(fileName)

	file, _ := os.Open(*fileName)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		splitText := strings.Fields(text)
		people = append(people, Person{fName: splitText[0], lName: splitText[1]})
	}

	for _, person := range people {
		fmt.Printf("%s %s\n", person.fName, person.lName)
	}
}