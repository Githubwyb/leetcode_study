package leetcode1

func TwoSum(nums []int, target int) (result []int) {
	hashMap := make(map[int]int)
	for i, v := range nums {
		expect := target - v
		index1, ok := hashMap[expect]
		if ok {
			result = append(result, index1)
			result = append(result, i)
			return
		}
		hashMap[v] = i
	}
	return
}
