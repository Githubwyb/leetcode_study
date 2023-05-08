package main

import (
	"container/list"
	. "leetcode/common"
)

func checkSymmetric(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	return a.Val == b.Val && checkSymmetric(a.Left, b.Right) && checkSymmetric(a.Right, b.Left)
}

func isSymmetric(root *TreeNode) bool {
	return checkSymmetric(root.Left, root.Right)
}

func isSymmetric1(root *TreeNode) bool {
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil || root.Right == nil {
		return false
	}
	var l list.List
	var r list.List
	l.PushBack(root.Left)
	r.PushBack(root.Right)
	for l.Len() != 0 {
		lt := l.Front().Value.(*TreeNode)
		l.Remove(l.Front())
		rt := r.Front().Value.(*TreeNode)
		r.Remove(r.Front())
		if lt == nil && rt == nil {
			continue
		}
		if (lt == nil || rt == nil) || lt.Val != rt.Val {
			return false
		}

		l.PushBack(lt.Left)
		r.PushBack(rt.Right)
		l.PushBack(lt.Right)
		r.PushBack(rt.Left)
	}
	return true
}
