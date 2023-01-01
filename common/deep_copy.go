package common

func DeepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}
