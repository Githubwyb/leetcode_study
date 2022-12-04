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

func removeNodes1(head *ListNode) *ListNode {
	arr := make([]*ListNode, 0, 1)
	for p := head; p != nil; p = p.Next {
		t := len(arr) - 1
		for t >= 0 && arr[t].Val < p.Val {
			arr = arr[:t]
			t--
		}
		arr = append(arr, p)
	}
	for i := 0; i < len(arr)-1; i++ {
		arr[i].Next = arr[i+1]
	}
	return arr[0]
}
