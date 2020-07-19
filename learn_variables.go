package main

import "fmt"

func learnvariablesMain() {

	/*Variable declaration going on here*/
	var z string /*Pre declaring variables*/
	var id int = 573
	name := "Sriraag"           /*shorthand variable declaration*/
	var x, y = "1234", false    /*declaring  multiple variables*/
	z = "Pre Declared Variable" /*assigning  values  for predeclared variables */

	/*Need to  read  up  on blank identifiers in go
	Go will not allow us to run a  program which has declared variables which  remain unused
	*/

	fmt.Println("Student-ID::", id)
	fmt.Println("Student-Name::", name)
	fmt.Println("x Value::", x)
	fmt.Println("y Value::", y)
	fmt.Println("z Value::", z)

}
