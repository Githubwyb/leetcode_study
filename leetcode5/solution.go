package main

import (
	_ "fmt"
)

func longestPalindrome(s string) (result string) {
	for i := range s {
		// think medium of palindrome as i
		left := i - 1
		right := i + 1
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			left--
			right++
		}
		left++
		if right-left > len(result) {
			result = s[left:right]
		}
		// think medium of palindrome as i and i+1
		left = i
		right = i + 1
		for left >= 0 && right < len(s) {
			if s[left] != s[right] {
				break
			}
			left--
			right++
		}
		left++
		if right-left > len(result) {
			result = s[left:right]
		}
	}
	return
}

func longestPalindrome1(s string) string {
	start, end := 0, -1
	t := "#"
	for i := 0; i < len(s); i++ {
		t += string(s[i]) + "#"
	}
	t += "#"
	s = t
	arm_len := []int{}
	right, j := -1, -1
	for i := 0; i < len(s); i++ {
		var cur_arm_len int
		if right >= i {
			i_sym := j*2 - i
			min_arm_len := min(arm_len[i_sym], right-i)
			cur_arm_len = expand(s, i-min_arm_len, i+min_arm_len)
		} else {
			cur_arm_len = expand(s, i, i)
		}
		arm_len = append(arm_len, cur_arm_len)
		if i+cur_arm_len > right {
			j = i
			right = i + cur_arm_len
		}
		if cur_arm_len*2+1 > end-start {
			start = i - cur_arm_len
			end = i + cur_arm_len
		}
	}
	ans := ""
	for i := start; i <= end; i++ {
		if s[i] != '#' {
			ans += string(s[i])
		}
	}
	return ans
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
