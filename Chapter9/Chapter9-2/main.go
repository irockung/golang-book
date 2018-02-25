package main

import (
	"fmt"
)

type Zipcode string

func main() {
	zipcode := Zipcode("11000")
	if zipcode.valid(){
		fmt.Println(zipcode)
	}
}


func (z Zipcode) valid() bool {
	if len(z)==5 {
		return true
	}
	return false
}