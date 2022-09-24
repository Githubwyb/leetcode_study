package main

import (
	_ "fmt"
)

type ProductOfNumbers1 struct {
	vec []int
	cnt []int
	pre []int
}

func Constructor1() (this ProductOfNumbers1) {
	return
}

func (this *ProductOfNumbers1) Add(num int) {
	n := len(this.vec)
	this.vec = append(this.vec, num)
	this.pre = append(this.pre, -1)

	// update cnt
	cnt := 0
	if n > 0 {
		// get last result
		cnt = this.cnt[n-1]
	}
	if num == 0 {
		cnt++
	}
	this.cnt = append(this.cnt, cnt)

	// update pre
	if n == 0 {
		return
	}
	if this.vec[n-1] <= 1 {
		this.pre[n] = this.pre[n-1]
	} else {
		this.pre[n] = n - 1
	}
}

func (this *ProductOfNumbers1) GetProduct(k int) int {
	n := len(this.vec)
	tot := this.cnt[n-1]
	if k < n {
		tot -= this.cnt[n-1-k]
	}
	if tot > 0 {
		return 0
	}

	result := 1
	// n-i-1 represent the next number to be used from the bottom
	// eg: i = n - 1, n-i-1 = 0, will use the 0th number from the bottom
	// which also represent the result contains 0 numbers
	for i := n - 1; (n - i - 1) < k; i = this.pre[i] {
		result *= this.vec[i]
	}
	return result
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
