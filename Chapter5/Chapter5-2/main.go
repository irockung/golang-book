package main

import (
	"fmt"
)

func f() int {
	return 1
}

func main() {

	i := 0
	if i == 0 {
		fmt.Println(i)
	} else {
		fmt.Println(i)
	}

	if i:= f(); i>0 {
		fmt.Println(i)
	}

	for number:= 1 ; number <= 100; number++ {
		if number%15 == 0 {
			fmt.Println(number, "FizzBuzz")
			continue
		}
		if number%3 == 0 {
			fmt.Println(number, "Fizz")
			continue
		}
		if number%5 == 0 {
			fmt.Println(number, "Buzz")
			continue
		}
		fmt.Println(number)
	}

	//i := num%2
	//switch i {
	//	case 0:fmt.Println("case", i)
	//	default:fmt.Println("default", i)
	//}
}

