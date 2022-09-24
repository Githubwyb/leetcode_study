package main

import (
	_ "fmt"
)

type ProductOfNumbers struct {
	result []int
}

func Constructor() (this ProductOfNumbers) {
	return
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.result = this.result[:0]
		return
	}
	if len(this.result) == 0 {
		this.result = append(this.result, num)
		return
	}
	this.result = append(this.result, this.result[len(this.result)-1]*num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k == len(this.result) {
		return this.result[len(this.result)-1]
	}
	if k > len(this.result) {
		return 0
	}
	return this.result[len(this.result)-1] / this.result[len(this.result)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
