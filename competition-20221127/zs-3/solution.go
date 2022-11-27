package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	cur := head
	arr := make([]int, 0, 1)
	for cur != nil {
		arr = append(arr, cur.Val)
		cur = cur.Next
	}
	max := arr[len(arr)-1]
	for i := len(arr) - 2; i >= 0; i-- {
		if max < arr[i] {
			max = arr[i]
		} else {
			arr[i] = max
		}
	}

	cur = head
	resultTmp := &ListNode{
		Next: head,
	}
	last := resultTmp
	for i := 0; i < len(arr); i++ {
		if cur.Val < arr[i] {
			last.Next = cur.Next
		} else {
			last = cur
		}
		cur = cur.Next
	}
	return resultTmp.Next
}
