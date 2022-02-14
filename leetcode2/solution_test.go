package leetcode2

import (
	"fmt"
	"testing"
)

func makeListNode(arr []int) (result *ListNode) {
	p := result
	for _, v := range arr {
		if p == nil {
			result = &ListNode{v, nil}
			p = result
			continue
		}
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	return
}

func TestTwoSum(t *testing.T) {
	type testCase struct {
		l1   *ListNode
		l2   *ListNode
		Want *ListNode
	}

	testGroup := []*testCase{
		{makeListNode([]int{0}), makeListNode([]int{0}), makeListNode([]int{0})},
		{makeListNode([]int{2, 4, 3}), makeListNode([]int{5., 6, 4}), makeListNode([]int{7, 0, 8})},
		{makeListNode([]int{9, 9, 9, 9, 9, 9, 9}), makeListNode([]int{9, 9, 9, 9}), makeListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}

	for i, v := range testGroup {
		result := AddTwoNumbers(v.l1, v.l2)
		p1 := result
		p2 := v.Want
		for p1 != nil || p2 != nil {
			if p1 == nil || p2 == nil {
				t.Fatalf("length not match, %d, l1 %s, l2 %s, expect %s but %s",
					i, PrintListNode(v.l1), PrintListNode(v.l2), PrintListNode(v.Want), PrintListNode(result))
			}
			if p1.Val != p2.Val {
				t.Fatalf("value not match, %d, l1 %s, l2 %s, expect %s but %s",
					i, PrintListNode(v.l1), PrintListNode(v.l2), PrintListNode(v.Want), PrintListNode(result))
			}
			p1 = p1.Next
			p2 = p2.Next
		}
		fmt.Println(i, PrintListNode(result))
	}
}
