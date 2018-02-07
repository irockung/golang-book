package main

import ("fmt")

func main() {
	fmt.Println("=====Zero Values=====")
	fmt.Println("0.1 + 0.2 ==  0.3 is", 0.1 + 0.2 == 0.3)
	num := 0.1
	num += 0.2
	fmt.Println("num == 0.3 is", num==0.3)
	fmt.Println("num is",num)
}