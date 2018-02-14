package main

import (
	"fmt"
)

func main() {
	for i, c := range "Golang" {
		fmt.Println(i, c)
		fmt.Printf("%v\n", string(c))
	}
}