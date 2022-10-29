package main

func longestContinuousSubstring(s string) int {
	result := 0
	curLen := 0
	for i := range s {
		if curLen == 0 {
			curLen = 1
			continue
		}

		if s[i]-s[i-1] == 1 {
			curLen++
			continue
		}
		if curLen > result {
			result = curLen
		}
		curLen = 1
	}
	if curLen > result {
		result = curLen
	}
	return result
}
