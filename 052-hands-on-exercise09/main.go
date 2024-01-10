package main

import "fmt"

// long declaration
var globalVar = 42

const constantVar = 420

func main() {
	//I did it this way
	//short declaration which infers the varaibles type
	shortDeclare := 4.20

	fmt.Println("My globalVar is: ", globalVar)
	fmt.Println("My global const var is: ", constantVar)
	fmt.Println("My short declaration var is: ", shortDeclare)

	//Todd did it this way
	//%v will show the value of a variable
	//%T will show the type of the variable
	fmt.Printf("the value of globalVar is %v and the type is %T\n", globalVar, globalVar)
	fmt.Printf("the value of constantVar is %v and the type is %T\n", constantVar, constantVar)
	fmt.Printf("the value of shortDeclare is %v and the type is %T\n", shortDeclare, shortDeclare)

}
