package main

import "fmt"

func main(){

//// Maps in Go

emp_details:=map[int]string{
1:"luffy",
2:"goku",
3:"uarmeshi",
4:"xyz",
}

fmt.Println(emp_details)

delete(emp_details,4)

fmt.Println(emp_details)

//// Another way to declare and assign values to map

emp_org:=map[int]string{}

emp_org[1]="one piece"
emp_org[2]="dragon ball"
emp_org[3]="yu yu hakushu"

fmt.Println(emp_org)

}

/* ++Output++
map[1:luffy 2:goku 3:uarmeshi 4:xyz]
map[1:luffy 2:goku 3:uarmeshi]
map[1:one piece 2:dragon ball 3:yu yu hakushu]
*/
