package main

import (
	"fmt"
	)
	
	
func main(){
fmt.Println("hello world") //Simple Hello World 

// Variable Declaration Examples


// +++++++++ Declaring and Separately Initializing Variables ++++++++++++
//Declaring a simple variable a of type int
//Then Separately initializing the variable
var a int;
a=10;

// +++++++++ Declaring/Initializing Variables together ++++++++++++
//Declaring and initializing a variable "name" of type string
// Here if the type is not mentioned it will be inferred from the value. of the variable
var name string="sriraag"


//Printing the declared/initialized variable
fmt.Printf("ID\t\t\t::%d\nName\t\t\t::%s \n",a,name)

// +++++++++ Multiple Variable Declaration ++++++++++++
// Another way to declare variable i.e mutliple variable declartion/initialization 

var x,y,z int= 1,2,3

fmt.Printf("Random_Numbers \t\t::%d %d %d\n",x,y,z)


// +++++++++ Short Variable Declaration ++++++++++++
address := "Bangalore, Karnataka India" 	// Here Type is automatically inferred from the value assigned
age,_ := 29,"M" 				// I can use "_" if i've assigned a value to a variable but dont want to use it yet since everything defined should be used in go

fmt.Printf("Address\t\t\t::%s\n",address)
///fmt.Printf("Age\t\t\t::%d\nGender\t\t\t::%s",age,gender)
fmt.Printf("Age\t\t\t::%d",age)
}




// Output
//hello world
//ID			::10
//Name			::sriraag 
//Random_Numbers 		::1 2 3
//Address			::Bangalore, Karnataka India
//Age			::29


