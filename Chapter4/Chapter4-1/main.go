package main

import (
	"fmt"
)

var (
	a = 1
	b = 2
)


var v1, v2 string = "1", "2"


func main() {

	v1, v2 := "1", "2"
	// Long use global
	var x string = "Hello, World"
	fmt.Println(x)
	fmt.Printf("Type : %T\n",x)

	var y string
	y = "Hello, World"
	fmt.Println(y)
	fmt.Printf("Type : %T\n",y)

	// Short not use global
	z := "Hello, World"
	fmt.Println(z)
	fmt.Printf("Type : %T\n",z)

	const c string = "unchange"
	fmt.Println(c)
	fmt.Printf("Type : %T\n",c)

	fmt.Println("v1", v1)
	fmt.Println("v2", v2)
}


