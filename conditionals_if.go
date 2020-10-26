package main

import "fmt"


func main(){

student:=map[int]string{}

student[1]="apple"
student[2]="apple"
fmt.Println(student)


	if student[1] == "sri" {
			fmt.Println("string matches")
	}else if student[2] == "apple" {
			fmt.Println("entered value is a fruit::",student[1])
	}else {
		fmt.Println("Unknown Value")
	}

var1:=24

if var1 <=30{

fmt.Println("Value is less than 30::",var1)
}else{

fmt.Println("Value is greater than 30::",var1)
}


}

/*


if conditional {

}else if conditional{

}else {

}
*/
