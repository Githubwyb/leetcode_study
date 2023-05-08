package common

import "sort"

func CompareSlice[T comparable](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	if len(l) == 0 {
		return true
	}
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}

func CompareSlices[T comparable](l, r [][]T) bool {
	if len(l) != len(r) {
		return false
	}
	if len(l) == 0 {
		return true
	}
	for i := range l {
		if !CompareSlice(l[i], r[i]) {
			return false
		}
	}
	return true
}

type compareType interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string
}

func CompareUnorderSlice[T compareType](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	if len(l) == 0 {
		return true
	}

	a := DeepCopy(l)
	b := DeepCopy(r)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return CompareSlice(a, b)
}
