package main

import (
	_ "fmt"
)

func reformatNumber(number string) (result string) {
	var addStr []byte
	for i := range number {
		if number[i] == ' ' || number[i] == '-' {
			continue
		}
		addStr = append(addStr, number[i])
		if len(addStr) > 4 {
			result += string(addStr[:3]) + "-"
			addStr = addStr[3:]
		}
	}
	if len(addStr) == 4 {
		result += string(addStr[:2]) + "-" + string(addStr[2:])
	} else {
		result += string(addStr)
	}
	return
}

func reformatNumber1(number string) string {
	addStr := make([]byte, 0, 4)
	result := make([]byte, 0, len(number))
	for i := range number {
		if number[i] == ' ' || number[i] == '-' {
			continue
		}
		addStr = append(addStr, number[i])
		if len(addStr) == 3 {
			addStr = append(addStr, '-')
			result = append(result, addStr...)
			addStr = addStr[:0]
		}
	}
	result = append(result, addStr...)
	if result[len(result)-1] == '-' {
		return string(result[:len(result)-1])
	} else if len(result) > 4 && result[len(result)-2] == '-' {
		result[len(result)-2] = result[len(result)-3]
		result[len(result)-3] = '-'
	}
	return string(result)
}
