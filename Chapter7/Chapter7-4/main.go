package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3}
	fmt.Println(slice)

	newSlice := make([]int, 2)
	fmt.Println(newSlice)
	copy(newSlice, slice)
	fmt.Printf("slice %v\n", slice)
	fmt.Printf("newSlice %v\n", newSlice)
}