package main

import "strings"

func haveConflict(event1 []string, event2 []string) bool {
	if strings.Compare(event1[0], event2[1]) > 0 {
		return false
	}

	if strings.Compare(event1[1], event2[0]) < 0 {
		return false
	}

	return true
}
