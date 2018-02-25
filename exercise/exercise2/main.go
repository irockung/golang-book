package main

import (
	"fmt"
)

var value int = 0

func insertCoin(coin string) {
	if coin == "T"  {value += 10}
	if coin == "F"  {value += 5}
	if coin == "TW" {value += 2}
	if coin == "O"  {value += 1}
}

func selectDrink(drink string, money int) int{
	if drink == "SD" {money-=18}
	if drink == "CC" {money-=12}
	if drink == "DW" {money-=7}
	return money
	
}

func main(){
	value = 0
	insertCoin("T")
	insertCoin("T")
	insertCoin("F")
	insertCoin("TW")
	insertCoin("O")
	var money int = value
	money = selectDrink("DW", money)
	fmt.Println("insert money", money)
	var coins string = ""
	for money>0 {
		if money >= 10 {
			money -= 10
			coins += ",T"
			continue
		}
		if money >= 5 {
			money -= 5
			coins += ",F"
			continue
		}
		if money >= 2 {
			money -= 2
			coins += ",TW"
			continue
		}
		if money >= 1 {
			money -= 1
			coins += ",O"
			continue
		}
	}
	fmt.Println("coin", coins)
}