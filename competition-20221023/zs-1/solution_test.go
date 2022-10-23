package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		event1 []string
		event2 []string
		Want bool
	}

	testGroup := []testCase{
		{[]string{"01:15", "02:00"}, []string{"02:00", "03:00"}, true},
		{[]string{"01:00", "02:00"}, []string{"01:20", "03:00"}, true},
		{[]string{"10:00", "11:00"}, []string{"14:00", "15:00"}, false},
	}

	for i, v := range testGroup {
		result := haveConflict(v.event1, v.event2)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
