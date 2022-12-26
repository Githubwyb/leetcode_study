package main

import "strings"

func getLast(str string) byte {
	return str[len(str)-1]
}

func isCircularSentence(sentence string) bool {
	arr := strings.Split(sentence, " ")
	if arr[0][0] != getLast(arr[len(arr)-1]) {
		return false
	}

	for i := 1; i < len(arr); i++ {
		if getLast(arr[i-1]) != arr[i][0] {
			return false
		}
	}
	return true
}

func isCircularSentence1(sentence string) bool {
	if sentence[0] != getLast(sentence) {
		return false
	}
	for i := range sentence {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return true
}
