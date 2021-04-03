package main

import (
	"fmt"
)

func main() {

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
