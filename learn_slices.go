package main

import "fmt"

func main(){

////Declaring Slices in Go


// Declaring a variable [var name string="sriraag" OR name:="sriraag" ]
// Declaring an Array [var names=[2] string {"luffy","zoro"} OR names:=[2]string {"luffy","zoro"} ]

// +++++++++++ Declaring a Slice in GO


// Longer Way
var names=[]string {};
names=append(names,"luffy","zoro");  /// append (what slice to append to,values)

fmt.Println(names)


// Short hand

days:=[]int {};
days=append(days,1,2,3,4,5,6,7)

fmt.Println(names,days);


///////++++++++++ Declaring a Slice using make function

action:=make([]string,2,5)   //// make is an built in function make(type of slice,len,capacity)

action[0]="add"
action[1]="update"

fmt.Println("Elements in Array",action)
fmt.Println("Length of Array",len(action))
fmt.Println("Capacity of Array",cap(action))

/// action[2]="delete".  uncommenting this gives an index out of range error, since length is 2 [panic: runtime error: index out of range [2] with length 2]

///Trying to append to the slice

action=append(action,"delete")
action=append(action,"delete")
action=append(action,"delete") 
action=append(action,"delete")  /// Adding the 6th element in slice of size 5 and it automatically doubles the capacity of the slice 


fmt.Println("Elements in Array",action)
fmt.Println("Length of Array",len(action))
fmt.Println("Capacity of Array",cap(action))

}


/*
[luffy zoro]
[luffy zoro] [1 2 3 4 5 6 7]
Elements in Array [add update]
Length of Array 2
Capacity of Array 5
Elements in Array [add update delete delete delete delete]
Length of Array 6
Capacity of Array 10
*/
