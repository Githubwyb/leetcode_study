package main

import (
	"container/list"
	_ "fmt"
)

const base = 997

type dataT struct {
	key   int
	value int
}

type MyHashMap1 struct {
	data []list.List
}

func Constructor1() (this MyHashMap1) {
	this.data = make([]list.List, base)
	return
}

func (this *MyHashMap1) Put(key int, value int) {
	h := key % base
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(dataT); et.key == key {
			e.Value = dataT{key, value}
			return
		}
	}
	this.data[h].PushBack(dataT{key, value})
}

func (this *MyHashMap1) Get(key int) int {
	h := key % base
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(dataT); et.key == key {
			return et.value
		}
	}
	return -1
}

func (this *MyHashMap1) Remove(key int) {
	h := key % base
	for e := this.data[h].Front(); e != nil; e = e.Next() {
		if et := e.Value.(dataT); et.key == key {
			this.data[h].Remove(e)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
