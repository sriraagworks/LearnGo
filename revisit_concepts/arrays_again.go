package main

import "fmt"

func main() {

	fmt.Println(`Program for doing the below
1.Different Types of Variable Declaration
2.Array
	2.1 Different Types of Declaring an Array
	2.2 CRUD Operations on an Array`)

	//**********************************VARIABLE DECLARATION EXAMPLES START HERE *******************************************
	// Variable Declaration Type1
	// Declaration and Initialization
	var var1 int = 1
	fmt.Println(var1)
	// Variable Declaration Type2
	// First Declaration and then initializing the variables
	var var2 string
	var2 = "This is variable two"
	fmt.Println(var2)
	// Variable Declaration Type3
	// Short Hand Declaration
	var3 := "this is short hand variable declaration"
	fmt.Println(var3)
	// Variable Declaration Type3
	// Declaring Multiple Variables together
	//a,_:=1,"variable which will not be used"
	var4, var5 := 573, "sriraag"
	fmt.Println(var4, var5)

	//**********************************VARIABLE DECLARATION EXAMPLES END HERE *******************************************

	//**********************************ARRAY EXAMPLES END HERE *******************************************

	fmt.Println("Array Declarations Begin Here")

	// Declaring First and Initializing Later
	var array1 [2]int
	array1[0] = 0
	array1[1] = 1
	fmt.Println(array1)

	// Declaring and Initializing At the Same Time
	var array2 = [2]string{"yes", "no"}
	fmt.Println(array2)

	var a3 = [2]int{}
	a3[1] = 100000
	fmt.Println(a3)

	// Short Hand Array Declaration
	cust_resp := [2]string{}
	cust_resp[0] = "Agree"
	cust_resp[1] = "Dis-Agree"
	fmt.Println(cust_resp)

	array4 := [5]int{11, 22, 33, 44, 55}
	fmt.Println(array4)

	fmt.Println("length of Array4 is ::", len(array4))
	///USING FOR TO PRINT ARRAY CONTENTS

	for i := 0; i < len(array4); i++ {

		fmt.Printf("Array-4 Index::%d Value::%d\n", i, array4[i])

	}

	for _, value := range array4 {

		fmt.Println(value)
	}

}
