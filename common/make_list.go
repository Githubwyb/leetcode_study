package main

type ListNode struct {
	Val  int
	Next *ListNode
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
	}
	return head
}

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
