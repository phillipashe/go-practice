package main

import (
	"fmt"
	"math"
)

// remove anything after the decimal point; print result
func main() {

	float := new(float64)

	fmt.Printf("Enter a number:")
	fmt.Scan(float)

	fmt.Println(math.Trunc(*float))

}
