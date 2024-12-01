package main

import "strings"

func isSubstringPresent(s string) bool {
	n := len(s)
	for i := n-1; i > 0; i-- {
		a := []byte{s[i], s[i-1]}
		if strings.Contains(s, string(a)) {
			return true
		}
	}
	return false
}
