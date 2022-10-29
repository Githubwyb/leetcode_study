package main

import "sort"

type NameSlice struct {
	names   []string
	heights []int
}

func (x NameSlice) Len() int           { return len(x.names) }
func (x NameSlice) Less(i, j int) bool { return x.heights[i] > x.heights[j] }
func (x NameSlice) Swap(i, j int) {
	x.heights[i], x.heights[j] = x.heights[j], x.heights[i]
	x.names[i], x.names[j] = x.names[j], x.names[i]
}

func sortPeople(names []string, heights []int) []string {
	tmp := NameSlice{
		names:   names,
		heights: heights,
	}
	sort.Sort(tmp)
	return tmp.names
}
