package common

func DeepCopy[T any](in []T) []T {
	result := make([]T, len(in))
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

func DeepCopySlices[T any](in [][]T) [][]T {
	result := make([][]T, len(in))
	for i := range result {
		result[i] = make([]T, len(in[i]))
		copy(result[i], in[i])
	}
	return result
}

func DeepCopyStrings(in []string) []string {
	result := make([]string, len(in))
	copy(result, in)
	return result
}
