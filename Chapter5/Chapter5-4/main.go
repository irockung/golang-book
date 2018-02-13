package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(100)

	for i := 1; i < 10; i++ {
		fmt.Println()
		if i>5 {
			fmt.Printf("เกินแล้ว")
			break
		}else{
			fmt.Print("Enter a number: ", i, " ")
			var input int
			fmt.Scanf("%d\n", &input)
			output := input

			if randomNum == output {
				fmt.Printf("เจอแล้ว", randomNum)
				break
			}
			if randomNum > output {
				fmt.Printf("น้อยไป")
				continue
			}
			if randomNum < output {
				fmt.Printf("มากไป")
				continue
			}
		}


	} 




}