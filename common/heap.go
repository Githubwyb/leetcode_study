package common

import "sort"

// 小根堆
type lhp struct{ sort.IntSlice }

func (l *lhp) Push(x interface{}) { l.IntSlice = append(l.IntSlice, x.(int)) }
func (l *lhp) Pop() interface{} {
	n := len(l.IntSlice)
	x := l.IntSlice[n-1]
	l.IntSlice = l.IntSlice[:n-1]
	return x
}

// 大根堆
type bhp struct{ sort.IntSlice }

func (l *bhp) Less(i, j int) bool { return l.IntSlice.Less(j, i) }
func (l *bhp) Push(x interface{}) { l.IntSlice = append(l.IntSlice, x.(int)) }
func (l *bhp) Pop() interface{} {
	n := len(l.IntSlice)
	x := l.IntSlice[n-1]
	l.IntSlice = l.IntSlice[:n-1]
	return x
}
