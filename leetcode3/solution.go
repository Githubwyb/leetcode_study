package leetcode2

import (
	_ "fmt"
)

func LengthOfLongestSubstring(s string) (result int) {
	hashMap := make(map[byte]int)
	left := 0
	for i, v := range s {
		i1, ok := hashMap[byte(v)]
		if !ok {
			hashMap[byte(v)] = i
			continue
		}
		if result < len(hashMap) {
			result = len(hashMap)
		}
		for j := left; j < i1+1; j++ {
			delete(hashMap, byte(s[j]))
		}
		left = i1 + 1
	}
	if result < len(hashMap) {
		result = len(hashMap)
	}

	return
}
