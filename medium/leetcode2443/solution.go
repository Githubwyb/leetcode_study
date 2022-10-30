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

func toInt(v bool) int {
	if v {
		return 1
	}
	return 0
}

// s数字的字符串形式
// l，r 左右指针
// pre 左侧是否被进位的标识
func existReverseSum(s string, l, r int, pre int) bool {
	if l == 0 && s[l] == '1' && s[r] == '0' {
		return false
	}
	suf := 0 // 右侧是否进位的标识
	// 181 l为0，r为2
	for l < r {
		// 存在进位标识，那么左侧高位应该为1x
		checkLeft := int(s[l] - '0')
		checkRight := int(s[r] - '0')

		l++
		r--

		if checkLeft-suf == checkRight && (pre == 0 || checkRight <= 8) {
			// 左侧存在进位，则右侧减掉进位不可能大于8，因为9+9最多到18，不可能到9
			pre = 0
			suf = 1
		} else if pre == 1 && checkLeft-1-suf == checkRight {
			suf = 1
			pre = 1
		} else if pre == 0 && checkLeft-suf == checkRight {
			suf = 0
			pre = 0
		} else if pre == 0 && checkLeft-1-suf == checkRight {
			suf = 0
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
