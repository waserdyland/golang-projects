package main

import "fmt"

func main() {
	x := 42
	y := 420
	z := 4.20

	fmt.Printf("the value of globalVar is %v and the type is %T\n", x, x)
	fmt.Printf("the value of constantVar is %v and the type is %T\n", y, y)
	fmt.Printf("the value of shortDeclare is %v and the type is %T\n", z, z)
}
