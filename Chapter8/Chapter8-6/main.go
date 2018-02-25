package main

import (
	"fmt"
)

func main() {
	m := map[int]int{1:1,2:2,3:3,}
	double(m)
	fmt.Printf("original pointer %p\n", m)
	fmt.Printf("original value %v\n", m)
}

func double(nums map[int]int){
	fmt.Printf("original pointer %p\n", nums)
	for index := 0; index < len(nums); index++ {
		nums[index] *= 2
	}
	fmt.Println("original value", nums)
}