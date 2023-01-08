package common

type BigHeap []int

func (h *BigHeap) Len() int { return len(*h) }

// less必须满足当Less(i, j)和Less(j, i)都为false，则两个索引对应的元素相等
// 为true，i向栈顶移动；为false，j向栈顶移动
func (h *BigHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *BigHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *BigHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
