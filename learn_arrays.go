package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Executed at [Nano]", time.Now().UnixNano())
	fmt.Println("Executed at [Normal]", time.Now())

	//1st type of declaring an Array
	id := [2]int{1, 2}
	name := [3]string{"sri", "goku"}
	name[2] = "xyz"
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println("Specifically Checking an Index of an Array 'name' =>", name[2])

	//2nd type of declaring an Array pre declaring the variable
	var day [2]string
	day[0] = "Saturday"
	day[1] = "Sunday"
	fmt.Println(day)
}
