package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() (res string) {
	for l != nil {
		res += fmt.Sprint(l.Val)
		l = l.Next
		if l != nil {
			res += " -> "
		}
	}
	return
}

func MakeList(list []int) (head *ListNode) {
	if len(list) == 0 {
		return
	}

	head = &ListNode{
		Val:  list[0],
		Next: nil,
	}
	pre := head
	for i := 1; i < len(list); i++ {
		pre.Next = &ListNode{
			Val:  list[i],
			Next: nil,
		}
		pre = pre.Next
	}
	return
}

func CompareList(l, r *ListNode) bool {
	for l != nil || r != nil {
		if l == nil || r == nil {
			return false
		}
		if l.Val != r.Val {
			return false
		}
		l = l.Next
		r = r.Next
	}
	return true
}

func List2slice(head *ListNode) []int {
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
