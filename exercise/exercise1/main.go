package main

import (
	"strconv"
	"fmt"
)

func main() {
	fmt.Println(weatherCelsius(25, "Bangkok few cloud"))
	fmt.Println(weatherCelsius(34, "Tak many"))
	fmt.Println(weatherCelsius(17, "Phuket rainy"))
	fmt.Println(weatherCelsius(9, "Chiang-mai cold"))
	fmt.Println(weather2(5, "test"))
}

func weatherCelsius(celsius int, desc string) string {

	var ret string

	switch celsius {
	case 25:
		ret = " _  _" + "\n" + " _||_ " + "\n" + "|_  _| c" + "\n" + desc
	case 34:
		ret = " _   " + "\n" + " _||_|" + "\n" + " _|  | c" + "\n" + desc
	case 17:
		ret = "    _" + "\n" + "  |  |" + "\n" + "  |  | c" + "\n" + desc
	default:
		ret = "x"
	}
	return ret
}

var nums = map[string][]string{
	"1": []string{"   ", "  |", "  |"},	
	"2": []string{" _ ", " _|", "|_ "},
	"3": []string{" _ ", " _|", " _|"},
	"4": []string{"   ", "|_|", "  |"},
	"5": []string{" _ ", "|_ ", " _|"},
	"6": []string{" _ ", "|_ ", "|_|"},
	"7": []string{" _ ", "  |", "  |"},
	"8": []string{" _ ", "|_|", "|_|"},
	"9": []string{" _ ", "|_|", " _|"},
	"0": []string{" _ ", "| |", "|_|"},
}

func weather2(celsius int, desc string) string {
	var ret string
	cel := strconv.Itoa(celsius)

	var str [][]string
	for _,ascii := range cel {
		numi:=string(ascii)
		switch numi {
		case "1":str=append(str,nums[numi])
		case "2":str=append(str,nums[numi])
		case "3":str=append(str,nums[numi])
		case "4":str=append(str,nums[numi])
		case "5":str=append(str,nums[numi])
		case "6":str=append(str,nums[numi])
		case "7":str=append(str,nums[numi])
		case "8":str=append(str,nums[numi])
		case "9":str=append(str,nums[numi])
		default :str=append(str,nums[numi])
	    }
	}

	for n:=0;n<3;n++ {
		for m:=0;m<len(cel);m++ {
			ret+=str[m][n]
		}
		ret+="\n"
	}
	ret+=desc

	return ret
}