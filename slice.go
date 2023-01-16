package main

import "fmt"
import "strconv"
import "sort"

func main() {
	var input = new(string) 
	var slice = make([]int, 3)
	length := 0
	fmt.Println("Enter any number.  Type \"X\" to exit.")
	for {
		fmt.Scan(input)
		if *input == "X" {
			fmt.Println("Exiting")
			break
		}
		// check the input is a number
		if newInt, err := strconv.Atoi(*input); err == nil {
			slice = append(slice, newInt)
			// if 1st, 2nd, or 3rd entry, need to replace a zero rather than append
			if length < 3 {
				// find a zero 
				for i, v := range slice {
					if v == 0 {
						slice = append(slice[:i], slice[i+1:]...)
						break;
					}
				}
			}
			length++
			sort.Ints(slice)
		}
		
		fmt.Println(slice)
	}
}