package main

import "sort"

type AVLNode struct {
	Value       int
	Height      int
	Count       int // 该节点的值出现的次数
	Left, Right *AVLNode
}

type AVLTree struct {
	n    int
	Root *AVLNode
}

// 获取节点的高度
func getHeight(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// 获取节点的平衡因子
func getBalance(node *AVLNode) int {
	if node == nil {
		return 0
	}
	return getHeight(node.Left) - getHeight(node.Right)
}

// 更新节点的高度
func updateHeight(node *AVLNode) {
	node.Height = max(getHeight(node.Left), getHeight(node.Right)) + 1
}

// 右旋转
func rightRotate(y *AVLNode) *AVLNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	updateHeight(y)
	updateHeight(x)

	return x
}

// 左旋转
func leftRotate(x *AVLNode) *AVLNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	updateHeight(x)
	updateHeight(y)

	return y
}

// 获取两个整数中的较大值
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 插入节点
func (tree *AVLTree) Insert(value int) {
	tree.n++
	tree.Root = insertNode(tree.Root, value)
}

func insertNode(node *AVLNode, value int) *AVLNode {
	if node == nil {
		return &AVLNode{Value: value, Height: 1, Count: 1}
	}

	if value < node.Value {
		node.Left = insertNode(node.Left, value)
	} else if value > node.Value {
		node.Right = insertNode(node.Right, value)
	} else {
		node.Count++
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Value {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Value {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	return node
}

// 统计比某个值大的元素数量
func (tree *AVLTree) CountGreaterThan(value int) int {
	return countGreaterThan(tree.Root, value)
}

func countGreaterThan(node *AVLNode, value int) int {
	if node == nil {
		return 0
	}

	if value < node.Value {
		return node.Count + countGreaterThan(node.Right, value) + countGreaterThan(node.Left, value)
	} else if value > node.Value {
		return countGreaterThan(node.Right, value)
	} else {
		return countGreaterThan(node.Right, value)
	}
}

type sortOther struct {
	sort.IntSlice
	arr []int
}

func (x sortOther) Swap(i, j int) {
	x.arr[i], x.arr[j] = x.arr[j], x.arr[i]
	x.IntSlice[i], x.IntSlice[j] = x.IntSlice[j], x.IntSlice[i]
}

func resultArray(nums []int) []int {
	if len(nums) <= 2 {
		return nums
	}
	mark := make([]int, len(nums))
	arr1 := AVLTree{}
	arr1.Insert(nums[0])
	mark[0] = 1
	arr2 := AVLTree{}
	arr2.Insert(nums[1])
	mark[1] = 2
	for i, v := range nums[2:] {
		if arr1.CountGreaterThan(v) > arr2.CountGreaterThan(v) {
			arr1.Insert(v)
			mark[i+2] = 1
		} else if arr1.CountGreaterThan(v) < arr2.CountGreaterThan(v) {
			arr2.Insert(v)
			mark[i+2] = 2
		} else if arr1.n <= arr2.n {
			arr1.Insert(v)
			mark[i+2] = 1
		} else {
			arr2.Insert(v)
			mark[i+2] = 2
		}
	}
	sort.Sort(sortOther{IntSlice: mark, arr: nums})
	return nums
}
