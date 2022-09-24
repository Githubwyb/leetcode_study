package main

import (
	_ "fmt"
)

func matchSubStr(s string, i int, p string, j int) bool {
	if j == len(p) {
		return i == len(s)
	}
	// check is first charactor match
	firstMatch := i < len(s) && (p[j] == '.' || s[i] == p[j])
	// when match 'x*'
	if j+1 < len(p) && p[j+1] == '*' {
		// match none and skip || match one but not skip
		return matchSubStr(s, i, p, j+2) || (firstMatch && matchSubStr(s, i+1, p, j))
	}

	// other, match charactor directly
	return firstMatch && matchSubStr(s, i+1, p, j+1)
}

func isMatch(s string, p string) bool {
	return matchSubStr(s, 0, p, 0)
}

func getStatus(s string, i int, p string, j int) bool {
	if j < 0 {
		return i < 0
	}
	if i < 0 {
		return p[j] == '*' && getStatus(s, i, p, j-2)
	}

	if p[j] != '*' {
		return (p[j] == '.' || s[i] == p[j]) && getStatus(s, i-1, p, j-1)
	}
	return getStatus(s, i, p, j-2) || ((p[j-1] == '.' || s[i] == p[j-1]) && getStatus(s, i-1, p, j))
}

func isMatch1(s string, p string) bool {
	return getStatus(s, len(s)-1, p, len(p)-1)
}
