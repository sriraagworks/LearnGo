package main

import "fmt"

func main(){

//// Maps in Go
  //variable_name:=map[key_var_type]value_type
emp_details:=map[int]string{
1:"luffy",
2:"goku",
3:"uarmeshi",
4:"xyz",
}

fmt.Println(emp_details)

delete(emp_details,4) //Deletes a Value from emp_details map

fmt.Println(emp_details)

}
