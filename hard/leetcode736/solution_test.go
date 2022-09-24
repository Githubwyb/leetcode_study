package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	type testCase struct {
		expression string
		Want       int
	}

	testGroup := []testCase{
		{"(let x 2 (mult x (let x 3 y 4 (add x y))))", 14},
		{"(let x 3 x 2 x)", 2},
		{"(let x 1 y 2 x (add x y) (add x y))", 5},
		{"(let a1 3 b2 (add a1 1) b2)", 4},
		{"(let x 2 (add (let x 3 (let x 4 x)) x))", 6},
		{"(mult 3 (add 2 3))", 15},
		{"(let x 7 -12)", -12},
		{"(let x -2 y x y)", -2},
		{"(let a -122 b 0 (add (add 1 -4) (mult a 66)))", -8055},
	}

	for i, v := range testGroup {
		result := evaluate(v.expression)
		if result != v.Want {
			t.Fatalf("%d, expression %v, expect %v but %v", i, v.expression, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}

func Test_parseBrackets(t *testing.T) {
	type testCase struct {
		expression string
		Want       []string
	}

	testGroup := []testCase{
		{"(let x 2 (mult x (let x 3 y 4 (add x y))))", []string{"let", "x", "2", "(mult x (let x 3 y 4 (add x y)))"}},
		{"((let x) 3 x 2 x)", []string{"(let x)", "3", "x", "2", "x"}},
	}

	for i, v := range testGroup {
		result := parseExpression(v.expression)
		for i := range result {
			if result[i] != v.Want[i] {
				t.Fatalf("%d, expression %v, expect %v but %v", i, v.expression, v.Want, result)
			}
		}
		if len(result) != len(v.Want) {
			t.Fatalf("%d, expression %v, expect %v but %v", i, v.expression, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}
