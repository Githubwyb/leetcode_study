package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		sentence string
		Want     bool
	}

	testGroup := []testCase{
		{"leetcode exercises sound delightful", true},
		{"eetcode", true},
		{"Leetcode is cool", false},
	}

	for i, v := range testGroup {
		result := isCircularSentence(v.sentence)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
	for i, v := range testGroup {
		result := isCircularSentence1(v.sentence)
		if result != v.Want {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
