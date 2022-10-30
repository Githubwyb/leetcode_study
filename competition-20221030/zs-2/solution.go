package main

import (
	"container/heap"
	"strings"
)

type Ids struct {
	id      string
	popular int
}

type BigHeap []Ids

func (h *BigHeap) Len() int { return len(*h) }

// less必须满足当Less(i, j)和Less(j, i)都为false，则两个索引对应的元素相等
// i > j 就是大根堆，i < j 就是小根堆
func (h *BigHeap) Less(i, j int) bool {
	if (*h)[i].popular > (*h)[j].popular {
		return true
	} else if (*h)[i].popular == (*h)[j].popular && strings.Compare((*h)[i].id, (*h)[j].id) < 0 {
		return true
	}
	return false
}
func (h *BigHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *BigHeap) Push(x interface{}) {
	*h = append(*h, x.(Ids))
}

func (h *BigHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type Author struct {
	ids     BigHeap
	popular int
}

func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	allAuthor := make(map[string]*Author)
	maxPopular := 0
	for i := range creators {
		if _, ok := allAuthor[creators[i]]; !ok {
			allAuthor[creators[i]] = &Author{
				make([]Ids, 0),
				0,
			}
		}
		tmp := allAuthor[creators[i]]
		heap.Push(&(tmp.ids), Ids{ids[i], views[i]})
		tmp.popular += views[i]
		if tmp.popular > maxPopular {
			maxPopular = tmp.popular
		}
	}

	result := make([][]string, 0)
	for k, v := range allAuthor {
		if v.popular == maxPopular {
			tmp := []string{k, v.ids[0].id}
			result = append(result, tmp)
		}
	}

	return result
}
