package main

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func isItPossible(word1 string, word2 string) bool {
	set1 := make(map[byte]int)
	set2 := make(map[byte]int)
	for _, t := range word1 {
		v := byte(t)
		if set1[v] > 0 {
			set1[v]++
		} else {
			set1[v] = 1
		}
	}
	for _, t := range word2 {
		v := byte(t)
		if set2[v] > 0 {
			set2[v]++
		} else {
			set2[v] = 1
		}
	}

	// 两个字符串中不同字符个数相等直接返回成功
	diff := abs(len(set1) - len(set2))
	more, less := set1, set2
	if len(set1) < len(set2) {
		more, less = set2, set1
	}
	if diff == 0 {
		// 多不变少不变，多的和少的有相同的
		//			    多的和少的各有一个只有一个但对方没有的
		// 多加少加		多的和少的各有一个有两个的
		for k := range more {
			if less[k] > 0 {
				return true
			}
			if more[k] >= 2 {
				for k1 := range less {
					if less[k1] >= 2 {
						return true
					}
				}
			} else if more[k] == 1 && less[k] == 0 {
				for k1 := range less {
					if less[k1] == 1 && more[k1] == 0 {
						return true
					}
				}
			}
		}
		return false
	}
	if diff == 1 {
		// 多减少不变，多的要替换的只有1个的并且替换来的是替换后已经有的
		//			  少的要替换的是有2个的并且替换来的是已经有的
		// 多减少动了一下还是没变，多的要替换的只有1个的并且替换来的是替换后已经有的
		//						 少的要替换的是有1个的并且替换来的是少的没有的
		// 多不变少加，多的要替换的是有2个的并且替换来的是已经有的
		// 			  少的要替换的是有2个的并且替换来的是没有的
		// 多变了一下但还是没变少加，多的要替换的是只有1个的并且是少的没有的，少的替换过来的是有两个的也是多的没有的
		for k := range less {
			if less[k] >= 2 && more[k] > 0 {
				for k1 := range more {
					if more[k1] == 1 && less[k1] > 0 && k1 != k {
						return true
					}
					if more[k1] >= 2 && less[k1] == 0 {
						return true
					}
				}
			}
			if less[k] >= 2 && more[k] == 0 {
				for k1 := range more {
					if more[k1] == 1 && less[k1] == 0 {
						return true
					}
				}
			}
			if less[k] == 1 && more[k] > 0 {
				for k1 := range more {
					if k != k1 && more[k1] == 1 && less[k1] == 0 {
						return true
					}
				}
			}
		}
		return false
	}
	if diff == 2 {
		// 多减少加，多的要替换的是只有1个的并且替换来的是替换后已经有的
		// 			少的要替换的是有2个的并且替换来的是没有的
		for k := range more {
			if more[k] == 1 && less[k] == 0 {
				for k1 := range less {
					if less[k1] >= 2 && more[k1] > 0 && k1 != k {
						return true
					}
				}
			}
		}
		return false
	}
	return false
}
