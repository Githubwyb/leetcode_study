package leetcode

import (
	_ "fmt"
)

func LengthOfLongestSubstring(s string) (result int) {
	hashMap := make(map[byte]int)
	i := 0
	for j, v := range s {
		index, ok := hashMap[byte(v)]
		if ok && index >= i {
			// 0 1 2 3
			// a b c a
			// i     j
			length := j - i
			if result < length {
				result = length
			}
			i = index + 1
		}
		hashMap[byte(v)] = j
	}
	// 0 1 2 3
	// a b c d
	// i
	length := len(s) - i
	if result < length {
		result = length
	}

	return
}

func LengthOfLongestSubstring1(s string) (result int) {
	var charMap [128]int
	i := 0
	for j, v := range s {
		index := charMap[byte(v)]
		if index > i {
			// 0 1 2 3
			// a b c a
			// i     j
			length := j - i
			if result < length {
				result = length
			}
			i = index
		}
		charMap[byte(v)] = j + 1
	}
	// 0 1 2 3
	// a b c d
	// i
	length := len(s) - i
	if result < length {
		result = length
	}

	return
}

func LengthOfLongestSubstring2(s string) (result int) {
	var charMap [128]int
	i := 0
	for j, v := range s {
		index := charMap[byte(v)]
		// 0 1 2 3 4
		// a b c b d
		// i     j
		// index = 2
		if i < index {
			i = index
		}
		length := j - i + 1
		if result < length {
			result = length
		}
		charMap[byte(v)] = j + 1
	}

	return
}
