package main

import "strings"

func haveConflict(event1 []string, event2 []string) bool {
	return strings.Compare(event1[0], event2[1]) <= 0 && strings.Compare(event1[1], event2[0]) >= 0
}
