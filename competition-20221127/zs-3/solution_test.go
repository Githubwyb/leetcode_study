package main

import (
	"fmt"
	"testing"
)

func list2slice(head *ListNode) []int {
	if head == nil {
		return []int{}
	}
	result := make([]int, 0, 1)
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func makeList(list []int) *ListNode {
	if len(list) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  list[0],
		Next: nil,
	}
	curr := head
	for i := 1; i < len(list); i++ {
		curr.Next = &ListNode{
			Val:  list[i],
			Next: nil,
		}
		curr = curr.Next
	}
	return head
}

func TestSolution(t *testing.T) {
	type testCase struct {
		list []int
		Want []int
	}

	testGroup := []testCase{
		{[]int{5, 2, 13, 3, 8}, []int{13, 8}},
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1}},
	}

	for i, v := range testGroup {
		tmp := removeNodes(makeList(v.list))
		result := list2slice(tmp)
		if !compareSlice(result, v.Want) {
			t.Fatalf("%d, v %v expect '%v' but '%v'", i, v, v.Want, result)
		}
		fmt.Println(i, "result", result)
	}
}

func compareSlice(l, r interface{}) bool {
	switch l.(type) {
	case []int:
		a := l.([]int)
		b := r.([]int)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case []string:
		a := l.([]string)
		b := r.([]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case [][]string:
		a := l.([][]string)
		b := r.([][]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if !compareSlice(a[i], b[i]) {
				return false
			}
		}

	default:
		panic("not implement")
	}

	return true
}
