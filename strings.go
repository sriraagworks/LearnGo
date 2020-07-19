package main

import "fmt"

func stringMain() {
	fmt.Println("Hello There this is printed using single line string literal which uses \"\" \n ")
	fmt.Println(`Hello There this is printed using raw string line string
literal which uses  and the text is multiline and prints characters as is \n `)
	fmt.Println('L') /*This is a rune literal can't have more than one character in here and uses '' */

}
