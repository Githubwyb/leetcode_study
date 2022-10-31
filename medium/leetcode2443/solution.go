package main

import "fmt"

func reverse(num int) int {
	result := 0
	for ; num != 0; num /= 10 {
		result *= 10
		result += num % 10
	}
	return result
}

func sumOfNumberAndReverse(num int) bool {
	for i := 0; i <= num; i++ {
		if i > num/2 && i%10 != 0 {
			continue
		}
		if i+reverse(i) == num {
			return true
		}
	}
	return false
}

// s数字的字符串形式
// l，r 左右指针
// pre 左侧是否被进位的标识
func existReverseSum(s string, l, r int, pre int) bool {
	// 左侧第一个为1，不考虑进位的情况，右侧必须为1，否则就是 100 = 050 + 050的情况
	if l == 0 && s[l] == '1' && s[r] == '0' {
		return false
	}
	suf := 0 // 右侧是否进位的标识
	for l < r {
		checkLeft := int(s[l] - '0')
		checkRight := int(s[r] - '0')
		l++
		r--

		// 存在进位，右侧最小到1
		if suf == 1 && checkRight == 0 {
			return false
		}
		// 本身要进位，两数相加最多到18，也就是右侧不可能为9（存在进位上面判断过不可能为0，那么就是不存在进位不可能为9）
		if pre == 1 && suf == 0 && checkRight == 9 {
			return false
		}

		if checkLeft == checkRight-suf {
			// 如果左侧存在进位，那么两个数最大加出来只能到18
			suf = pre
			pre = 0
		} else if checkLeft-1+suf == checkRight {
			suf = pre
			pre = 1
		} else {
			return false
		}
	}
	if l == r && (int(s[l]-'0')-suf)%2 == 1 {
		// 13121 中间剩个奇数，就找不到
		return false
	} else if l != r && suf != pre {
		// 21 剩余一个suf没人认领
		return false
	}
	return true
}

func sumOfNumberAndReverse1(num int) bool {
	s := fmt.Sprint(num)
	n := len(s)

	return existReverseSum(s, 0, n-1, 0) || (s[0] == '1' && n > 1 && existReverseSum(s, 1, n-1, 1))
}
