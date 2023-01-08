package common

func DeepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}

func DeepCopyStrings(in []string) []string {
	result := make([]string, len(in))
	copy(result, in)
	return result
}
