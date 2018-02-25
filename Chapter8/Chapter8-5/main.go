package main

import (
	"fmt"
)

func main() {
	slice := []int{1,2,3}
	double(slice)
	fmt.Printf("original pointer %p\n", slice)
	fmt.Printf("original value %v\n", slice)
}

func double(nums []int){
	fmt.Printf("original pointer %p\n", nums)
	for index := 0; index < len(nums); index++ {
		nums[index] *= 2
	}
	fmt.Println("original value", nums)
}