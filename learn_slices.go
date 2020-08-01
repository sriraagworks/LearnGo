package main

import (
	"fmt"
)

func main() {

	//Normal Array Declaration
	id := [3]int{1, 2, 3}
	fmt.Println("Array Declaration", id)

	//1st Type of Slice Declaration

	//String Slice
	fruits := []string{}
	fruits = append(fruits, "apple", "mango", "banana")
	fmt.Println("String Slice =>", fruits)

	//Integer Slice
	year := []int{}
	year = append(year, 1992, 1993, 1994, 2007, 2017, 2018, 2020)

	fmt.Println("Integer Slice =>", year)

	//Creating a slice but limiting it with a max value beyond which the array wont grow
	ships := make([]string, 1)
	ships[0] = "a"
	fmt.Println("Using Make to set max limit for a slice declaration=>", ships)
}
