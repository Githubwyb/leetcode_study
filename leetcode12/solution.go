package main

import (
	_ "fmt"
)

var intA = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
var strA = [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

func intToRoman(num int) string {
	for i := len(intA) - 1; i >= 0; i-- {
		if num >= intA[i] {
			return strA[i] + intToRoman(num-intA[i])
		}
	}
	return ""
}

func intToRoman1(num int) (result string) {
	var intA = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var strA = [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}

	for i := len(intA) - 1; i >= 0; i-- {
		for num >= intA[i] {
			num -= intA[i]
			result += strA[i]
		}
	}
	return
}
