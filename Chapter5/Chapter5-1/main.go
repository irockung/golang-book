package main

import (
	"fmt"
)

func main() {
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")


	for i := 0; i < 10; i++ {
		fmt.Println("for simple", i)
	}

	j := 0
	for  {
		j++
		fmt.Println("while", j)
		if j == 10 {
			break
		}
	}



}