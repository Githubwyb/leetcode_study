package main

import (
	_ "fmt"
)

type MyHashMap struct {
	data []int
}

func Constructor() (this MyHashMap) {
	this.data = make([]int, 1e5)
	for i := range this.data {
		this.data[i] = -1
	}
	return
}

func (this *MyHashMap) Put(key int, value int) {
	this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
	return this.data[key]
}

func (this *MyHashMap) Remove(key int) {
	this.data[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
