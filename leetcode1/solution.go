package leetcode

func twoSum(nums []int, target int) []int {
	hashMap := make(map[int]int)
	for i, v := range nums {
		expect := target - v
		index1, ok := hashMap[expect]
		if ok {
			return []int{index1, i}
		}
		hashMap[v] = i
	}
	return nil
}
