package main

import (
	"fmt"
	"testing"
)

type testCase struct {
	op    []string
	param []int
	Want  []interface{}
}

func TestSolution(t *testing.T) {
	testGroup := []testCase{
		{
			[]string{"MinStack", "push", "push", "push", "top", "pop", "getMin", "pop", "getMin", "pop", "push", "top", "getMin", "push", "top", "getMin", "pop", "getMin"},
			[]int{0, 2147483646, 2147483646, 2147483647, 0, 0, 0, 0, 0, 0, 2147483647, 0, 0, -2147483648, 0, 0, 0, 0},
			[]interface{}{nil, nil, nil, nil, 2147483647, nil, 2147483646, nil, 2147483646, nil, nil, 2147483647, 2147483647, nil, -2147483648, -2147483648, nil, 2147483647},
		},
		{
			[]string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			[]int{0, -2, 0, -3, 0, 0, 0, 0},
			[]interface{}{nil, nil, nil, nil, -3, nil, 0, -2},
		},
	}

	for i, v := range testGroup {
		var obj MinStack
		for j := range v.op {
			if v.op[j] == "MinStack" {
				obj = Constructor()
			} else {
				handle(t, obj, v, j, i)
			}
		}
	}
	for i, v := range testGroup {
		var obj MinStack
		for j := range v.op {
			if v.op[j] == "MinStack" {
				obj = Constructor2()
			} else {
				handle(t, obj, v, j, i)
			}
		}
	}
	for i, v := range testGroup {
		var obj MinStack
		for j := range v.op {
			if v.op[j] == "MinStack" {
				obj = Constructor3()
			} else {
				handle(t, obj, v, j, i)
			}
		}
	}
}

func handle(t *testing.T, obj MinStack, v testCase, j int, i int) {
	switch v.op[j] {
	case "push":
		obj.Push(v.param[j])
	case "getMin":
		result := obj.GetMin()
		if result != v.Want[j].(int) {
			t.Fatalf("%d, v %v expect '%v' but %d '%v'", i, v, v.Want, j, result)
		}
		fmt.Println(i, "result", result)
	case "pop":
		obj.Pop()
	case "top":
		result := obj.Top()
		if result != v.Want[j].(int) {
			t.Fatalf("%d, v %v expect '%v' but %d '%v'", i, v, v.Want, j, result)
		}
		fmt.Println(i, "result", result)
	default:
		panic("not implement")
	}
}
