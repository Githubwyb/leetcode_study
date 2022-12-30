package main

import (
	"container/list"
	"fmt"
)

type mem struct {
	start int
	end   int // end不属于范围
	mID   int
}

type Allocator struct {
	memMap  *list.List
	freeMap map[int][]*list.Element
}

func Constructor(n int) Allocator {
	res := Allocator{}
	res.memMap = list.New()
	// 放一个头节点进去
	res.memMap.PushFront(mem{
		start: 0,
		end:   0,
		mID:   -1,
	})
	// 放一个尾节点进去
	res.memMap.PushBack(mem{
		start: n,
		end:   n,
		mID:   -1,
	})
	res.freeMap = make(map[int][]*list.Element)
	fmt.Printf("init %p\r\n", &(res.memMap))
	return res
}

func (this *Allocator) Allocate(size int, mID int) int {
	// 从第一个开始找所有间隔是否有空位
	p := this.memMap.Front()
	for n := p.Next(); n != nil; p, n = n, n.Next() {
		if n.Value.(mem).start-p.Value.(mem).end < size {
			continue
		}
		start := p.Value.(mem).end
		fmt.Printf("insert %p\r\n", &(this.memMap))
		this.memMap.InsertAfter(mem{
			start: start,
			end:   start + size,
			mID:   mID,
		}, p)
		if _, ok := this.freeMap[mID]; !ok {
			this.freeMap[mID] = make([]*list.Element, 0, 1)
		}
		this.freeMap[mID] = append(this.freeMap[mID], p.Next())
		return start
	}
	return -1
}

func (this *Allocator) Free(mID int) int {
	ans := 0
	if idMap, ok := this.freeMap[mID]; ok {
		for _, v := range idMap {
			ans += v.Value.(mem).end - v.Value.(mem).start
			this.memMap.Remove(v)
		}
		this.freeMap[mID] = this.freeMap[mID][:0]
	}
	return ans
}

/**
 * Your Allocator object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Allocate(size,mID);
 * param_2 := obj.Free(mID);
 */
