package leetcode

import (
	_ "fmt"
)

// s1 range [0, m]
// s2 range [0, n]
// k range [1, (m+n+1)/2]
func getKth(nums1 []int, s1 int, nums2 []int, s2 int, k int) int {
	m := len(nums1)
	n := len(nums2)

	if s1 == m {
		return nums2[s2+k-1]
	}
	if s2 == n {
		return nums1[s1+k-1]
	}
	if k == 1 {
		if nums1[s1+k-1] > nums2[s2+k-1] {
			return nums2[s2+k-1]
		}
		return nums1[s1+k-1]
	}
	index := k / 2
	var a1 int
	var a2 int
	a1 = s1 + index - 1
	if a1 >= m {
		a1 = m - 1
	}
	a2 = s2 + index - 1
	if a2 >= n {
		a2 = n - 1
	}
	if nums1[a1] > nums2[a2] {
		return getKth(nums1, s1, nums2, a2+1, k+s2-a2-1)
	} else {
		return getKth(nums1, a1+1, nums2, s2, k+s1-a1-1)
	}
}

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	left := getKth(nums1, 0, nums2, 0, (m+n+1)/2)
	if (m+n)%2 == 1 {
		return float64(left)
	}

	right := getKth(nums1, 0, nums2, 0, (m+n)/2+1)
	return float64(left+right) / 2
}

func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	if m > n {
		return FindMedianSortedArrays1(nums2, nums1)
	}

	i := 0    // value range: [0, m]
	var j int // j = (m+n+1) / 2 - i, value range: [(m-n) / 2, (m+n+1) / 2] => [0, n]
	left := 0
	right := m
	for {
		// 0, 1, 2 => 1
		// 0, 1 => 0
		i = (left + right) / 2
		// i == 0 => j = (2+3) / 2 - 0 = 2:
		// 0 1   |    0 1 2    => left   b0 b1
		// i     |        j    => right  a0 a1 b2
		j = (m+n)/2 - i

		// m == 0 => i == 0, so (i != 0 || i != m) => m != 0
		if i != 0 && j != n && nums1[i-1] > nums2[j] {
			// i is bigger than expect, so i != 0 && j != m
			// index | 0 1 2 3 |  0 1 2 3 4
			// value | 1 3 5   |  2 4 6 8 10
			//       |       i |    j
			right = i - 1
		} else if i != m && j != 0 && nums2[j-1] > nums1[i] {
			// i is less than expect, so i != m && j != 0
			// index | 0 1 2 3 |  0 1 2 3 4
			// value | 1 3 5   |  2 4 6 8 10
			//       | i       |          j
			left = i + 1
		} else {
			// m == 0, right must be nums[j]
			// i == m, right must be nums[j]
			// j == n, right must be nums[i]
			var rightMin int
			if m == 0 || i == m || (j != n && nums1[i] > nums2[j]) {
				rightMin = nums2[j]
			} else {
				rightMin = nums1[i]
			}
			if (m+n)%2 == 1 {
				return float64(rightMin)
			}

			// m == 0, left must be nums[j-1]
			// i == 0, left must be nums[j-1]
			// j == 0, left must be nums[i-1]
			var leftMax int
			if m == 0 || i == 0 || (j != 0 && nums1[i-1] < nums2[j-1]) {
				leftMax = nums2[j-1]
			} else {
				leftMax = nums1[i-1]
			}

			return float64(leftMax+rightMin) / 2
		}
	}
}

// s1 range [0, m]
// s2 range [0, n]
// k range [1, (m+n+1)/2]
func findMedianOfTwoArray(nums1 []int, s1 int, nums2 []int, s2 int, k int) float64 {
	m := len(nums1)
	n := len(nums2)
	// s1 == m, nums2[s2+k-1] & nums2[s2+k]
	// s2 == n, nums1[s1+k-1] & nums1[s1+k]
	// k == 1, min(nums1[s1], nums2[s2])

	if s1 == m {
		if (m+n)%2 == 1 {
			return float64(nums2[s2+k-1])
		}
		return float64(nums2[s2+k-1]+nums2[s2+k]) / 2
	}
	if s2 == n {
		if (m+n)%2 == 1 {
			return float64(nums1[s1+k-1])
		}
		return float64(nums1[s1+k-1]+nums1[s1+k]) / 2
	}

	if k == 1 {
		var left int
		if nums1[s1] > nums2[s2] {
			left = nums2[s2]
			s2++
		} else {
			left = nums1[s1]
			s1++
		}
		if (m+n)%2 == 1 {
			return float64(left)
		}

		var right int
		if s1 == m {
			right = nums2[s2]
		} else if s2 == n {
			right = nums1[s1]
		} else if nums1[s1] > nums2[s2] {
			right = nums2[s2]
		} else {
			right = nums1[s1]
		}
		return float64(left+right) / 2
	}
	index := k / 2
	var a1 int
	var a2 int
	a1 = s1 + index - 1
	if a1 >= m {
		a1 = m - 1
	}
	a2 = s2 + index - 1
	if a2 >= n {
		a2 = n - 1
	}
	if nums1[a1] > nums2[a2] {
		return findMedianOfTwoArray(nums1, s1, nums2, a2+1, k+s2-a2-1)
	} else {
		return findMedianOfTwoArray(nums1, a1+1, nums2, s2, k+s1-a1-1)
	}
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	return findMedianOfTwoArray(nums1, 0, nums2, 0, (len(nums1)+len(nums2)+1)/2)
}
