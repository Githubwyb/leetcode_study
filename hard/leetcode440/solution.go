package main

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

// 获取当前节点的子节点个数，n为总数
func getChildCount(cur int32, n int32) int64 {
	// 假设cur为1，第一层为10-19
	//            第二层100-199
	//            第三层1000-1999
	var count int64 = 0
	var start int64 = int64(cur * 10)
	var end int64 = start + 9
	for start <= int64(n) {
		// n为13，那么这一层就是13-10+1=4个节点
		count += min(end, int64(n)) - start + 1
		start *= 10
		end = end*10 + 9
	}
	return count
}

func findKthNumber(n int32, k int32) int32 {
	var cur int32 = 1
	k-- // k为0时，cur就是所求。如果k为1，那么返回第一个也就是1，所以这里减一
	for k > 0 {
		// 获取当前节点的子节点个数
		count := getChildCount(cur, n)
		// 当k大于当前节点的子节点数，假设n为11，k为4。 1 10 11 2 3 4 5 6 7 8 9
		// 到这里k为3，count为2，计算后k为0，cur为2
		if int64(k) > count {
			k -= int32(count + 1)
			cur++
			continue
		}
		// k小于等于count
		// 小于的情况，假设n为11，k为2。 1 10 11 2 3 4 5 6 7 8 9
		//     到这里k为1，count为2，计算后k为0，cur为10
		// 等于的情况，假设k为3
		//     到这里k为2，计算后k为1，cur为10
		cur *= 10
		k--
	}
	return cur
}
