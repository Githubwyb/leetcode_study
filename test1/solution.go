package main

import "sort"

// https://leetcode.cn/problems/longest-word-lcci/

type sortWord []string

func (s *sortWord) Len() int { return len(*s) }
func (s *sortWord) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}
func (s *sortWord) Less(i, j int) bool {
	if len((*s)[i]) == len((*s)[j]) {
		return (*s)[i] > (*s)[j]
	}
	return len((*s)[i]) > len((*s)[j])
}

func ComboWords(in []string) []string {
	// 对in进行长度的从大到小排序
	sort.Sort(&sortWord{})
	// 使用map储存所有in的单词，方便查找
	// 将最长的取出来向前查找是否存在可以完全匹配前缀的子串，找到就剪掉（剪掉先查看map是否有直接匹配的）
	// 没有就继续查找（从长度小于裁剪后的结果的地方向前查找），可能存在多个完全匹配前缀的子串，都遍历一遍
	// 找到最后没有找到就跳过从下一个开始查找
}
