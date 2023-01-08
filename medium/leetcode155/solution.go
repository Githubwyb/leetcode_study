package main

import "math"

type MinStack interface {
	Push(val int)
	Pop()
	Top() int
	GetMin() int
}

type MinStack1 struct {
	stack []int
	min   int
}

func Constructor() MinStack {
	return &MinStack1{
		stack: make([]int, 0),
	}
}

func (this *MinStack1) Push(val int) {
	if len(this.stack) == 0 {
		this.min = val
		this.stack = append(this.stack, val-this.min)
	} else {
		this.stack = append(this.stack, val-this.min)
		if val < this.min {
			this.min = val
		}
	}
}

func (this *MinStack1) Pop() {
	if len(this.stack) == 0 {
		return
	}
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if v < 0 {
		this.min -= v
	}
}

func (this *MinStack1) Top() int {
	v := this.stack[len(this.stack)-1]
	if v < 0 {
		return this.min
	}
	return this.min + v
}

func (this *MinStack1) GetMin() int {
	return this.min
}

type MinStack2 struct {
	stack []int
	min   int
}

func Constructor2() MinStack {
	return &MinStack2{
		stack: make([]int, 0),
		min:   math.MaxInt,
	}
}

func (this *MinStack2) Push(val int) {
	if val <= this.min {
		this.stack = append(this.stack, this.min)
		this.min = val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack2) Pop() {
	if len(this.stack) == 0 {
		return
	}
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if v == this.min {
		this.min = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack2) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack2) GetMin() int {
	return this.min
}

type MinStack3 struct {
	stack    []int
	minStack []int
}

func Constructor3() MinStack {
	return &MinStack3{
		stack:    make([]int, 0),
		minStack: make([]int, 0),
	}
}

func (this *MinStack3) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack3) Pop() {
	if len(this.stack) == 0 {
		return
	}
	v := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if this.minStack[len(this.minStack)-1] == v {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack3) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack3) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
