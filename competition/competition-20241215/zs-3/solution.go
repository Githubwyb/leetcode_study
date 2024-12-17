package main

func z_func(nums []int) []int {
	n := len(nums)
	z := make([]int, n)
	zl, zr := 0, 0
	for i := 1; i < n; i++ {
		if i < zr && z[i-zl] < zr-i {
			// z[i-l]对应的数组没有超出zbox
			z[i] = z[i-zl]
			continue
		}
		// i超出zbox，或长度超出zbox重新算
		if i < zr {
			// i在zbox中，从zr开始找，计算少一点
			z[i] = zr - i
		} else {
			// i不在zbox中，从i开始找zr
			zr = i
		}
		for ; zr < n && nums[z[i]] == nums[zr]; zl, zr, z[i] = i, zr+1, z[i]+1 {
		}
	}
	return z
}

func beautifulSplits(nums []int) int {
	n := len(nums)
	z := z_func(nums)
	res := 0
	for i := 1; i < n-1; i++ {
		z0 := z_func(nums[i:])
		for j := i + 1; j < n; j++ {
			// nums1是nums2的前缀，len(nums1) <= len(nums2) && z[i] >= len(nums1)
			// nums2是nums3的前缀，len(nums2) <= len(nums3) && z0[j-i] >= len(nums2)
			if i <= j-i && z[i] >= i {
				res++
			} else if j-i <= n-j {
				if z0[j-i] >= j-i {
					res++
				}
			}
		}
	}
	return res
}

func beautifulSplits1(nums []int) int {
	n := len(nums)
	// 多一个，最后一个越界的n是0，不影响计算
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j > i; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	res := 0
	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			// nums1是nums2的前缀，len(nums1) <= len(nums2) && lcp[nums1][nums2] >= len(nums1)
			// nums2是nums3的前缀，len(nums2) <= len(nums3) && lcp[nums2][nums3] >= len(nums2)
			if (i <= j-i && lcp[0][i] >= i) || (j-i <= n-j && lcp[i][j] >= j-i) {
				res++
			}
		}
	}

	return res
}
