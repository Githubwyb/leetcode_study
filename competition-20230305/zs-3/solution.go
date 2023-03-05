package main

const mx int = 1e6

var primes []int = make([]int, 0, 78498)

func getAllPrime() {
	flag := make([]bool, mx+1) // 标记数有没有被筛掉，false就是没有
	for i := 2; i < mx+1; i++ {
		if !flag[i] {
			// 数没有被比自己小的数筛掉，就代表是质数
			primes = append(primes, i)
		}
		for _, v := range primes {
			if i*v > mx {
				break
			}
			// 每一个数都作为因子乘以比自己小的素数筛掉后面的数
			flag[i*v] = true
			if i%v == 0 {
				// 减少时间复杂度的关键算法
				// 12 = 2 * 3 * 2，i = 4时，只排了8就退出了，因为6会将12排除
				// 也就是，假设v可以整除i即i = kv，有某个数为x = mi = kmv
				//        那么存在一个数 i < km < x可以把x排掉，用i乘以所有的质数去排除就没什么意义了，提前退出减少时间复杂度
				break
			}
		}
	}
}

func init() {
	getAllPrime()
}

func findValidSplit(nums []int) int {
	check := make([]int, len(primes))
	for _, v := range nums {
		for i := 0; i < len(primes); i++ {
			if v%primes[i] == 0 {
				check[primes[i]]++
				v /= primes[i]
				i = -1
			}
		}
	}
	target := make(map[int]int)
	for i := range nums {
		
	}
}
