package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		num  int
		Want string
	}

	testGroup := []*testCase{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for i, v := range testGroup {
		result := intToRoman(v.num)
		if result != v.Want {
			t.Fatalf("%d, num %v, expect %v but %v", i, v.num, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
	for i, v := range testGroup {
		result := intToRoman1(v.num)
		if result != v.Want {
			t.Fatalf("%d, num %v, expect %v but %v", i, v.num, v.Want, result)
		}
		fmt.Println(i, "result:", result)
	}
}
