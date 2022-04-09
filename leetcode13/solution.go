package main

import (
	_ "fmt"
)

func romanToInt(s string) (result int) {
	var intA = [...]int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	var strA = [...]string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	index := 0
	for i := len(strA) - 1; i >= 0; i-- {
		compareLen := len(strA[i])
		for index+compareLen <= len(s) && s[index:index+compareLen] == strA[i] {
			result += intA[i]
			index += compareLen
		}
	}
	return
}

func romanToInt1(s string) (result int) {
	romToIntMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := range s {
		value := romToIntMap[s[i]]
		if i < len(s)-1 && romToIntMap[s[i+1]] > value {
			result -= value
			continue
		}
		result += value
	}
	return
}
