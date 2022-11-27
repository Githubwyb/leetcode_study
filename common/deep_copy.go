package main

func deepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}
