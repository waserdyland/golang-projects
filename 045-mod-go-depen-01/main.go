package main

import (
	"fmt"

	//gets puppy.go straight from github
	"github.com/waserdyland/puppy"
)

func main() {
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	fmt.Println(s1)
	fmt.Println(s2)

	//all in one
	//fmt.Println(puppy.Bark())
	//fmt.Println(puppy.Barks())

}
