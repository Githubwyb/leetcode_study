package leetcode

import (
	. "leetcode/common"
)

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	flag := 0
	p := result
	for l1 != nil || l2 != nil || flag != 0 {
		val := flag
		if l1 != nil {
			val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			val += l2.Val
			l2 = l2.Next
		}
		var tmp ListNode
		tmp.Val, flag = val%10, val/10
		if p == nil {
			result = &tmp
			p = result
			continue
		}
		p.Next = &tmp
		p = p.Next
	}
	return
}
