package common

type IntsSlice [][]int

func (x IntsSlice) Len() int { return len(x) }

// 为true，i向前；false，j向前。要满足相等时返回false
func (x IntsSlice) Less(i, j int) bool { return x[i][0] < x[j][0] }
func (x IntsSlice) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
