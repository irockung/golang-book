package main

import (
	"fmt"
)

func Fizz(number int) string {
	if number%3 == 0 {
		return "Fizz"
	}
	return ""
}

func Buzz(number int) string {
	if number%5 == 0 {
		return "Buzz"
	}
	return ""
}

func main() {

	for number:= 1 ; number <= 100; number++ {
		fmt.Println(number, Fizz(number)+Buzz(number))
	}

}

