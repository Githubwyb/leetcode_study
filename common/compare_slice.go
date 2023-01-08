package common

import "sort"

func CompareSlice(l, r interface{}) bool {
	switch l.(type) {
	case []int:
		a := l.([]int)
		b := r.([]int)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case []string:
		a := l.([]string)
		b := r.([]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if a[i] != b[i] {
				return false
			}
		}

	case [][]string:
		a := l.([][]string)
		b := r.([][]string)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if !CompareSlice(a[i], b[i]) {
				return false
			}
		}
	case [][]int:
		a := l.([][]int)
		b := r.([][]int)
		if len(a) != len(b) {
			return false
		}

		if len(a) == 0 {
			return true
		}

		for i := range a {
			if !CompareSlice(a[i], b[i]) {
				return false
			}
		}

	default:
		panic("not implement")
	}

	return true
}

func CompareUnorderSlice(l, r interface{}) bool {
	switch l.(type) {
	case []int:
		a := DeepCopy(l.([]int))
		b := DeepCopy(r.([]int))
		sort.Ints(a)
		sort.Ints(b)
		return CompareSlice(a, b)

	case []string:
		a := DeepCopyStrings(l.([]string))
		b := DeepCopyStrings(r.([]string))
		sort.Strings(a)
		sort.Strings(b)
		return CompareSlice(a, b)

	default:
		panic("not implement")
	}
}
