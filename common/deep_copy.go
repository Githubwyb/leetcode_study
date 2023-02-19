package common

func DeepCopy(in []int) []int {
	result := make([]int, len(in))
	copy(result, in)
	return result
}

func DeepCopyIntss(in [][]int) [][]int {
	result := make([][]int, len(in))
	for i := range result {
		result[i] = make([]int, len(in[i]))
		copy(result[i], in[i])
	}
	return result
}

func DeepCopyStrings(in []string) []string {
	result := make([]string, len(in))
	copy(result, in)
	return result
}
