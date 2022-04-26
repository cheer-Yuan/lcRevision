package main

import "fmt"

/* 接雨水
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6

当前列雨水面积：min(左边柱子的最高高度，右边柱子的最高高度) - 当前柱子高度。
递推公式：
maxLeft[i] = max(height[i], maxLeft[i - 1])
maxRight[i] = max(height[i], maxRight[i + 1])
*/

func trap(height []int) int {
	maxLeft, maxRight, sum := make([]int, len(height)), make([]int, len(height)), 0
	maxLeft[0], maxRight[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		maxLeft[i] = MaxOf2(height[i], maxLeft[i-1])
		maxRight[len(height)-1-i] = MaxOf2(height[len(height)-1-i], maxRight[len(height)-i])
	}

	for i := 1; i < len(maxLeft)-1; i++ {
		buff := MinOf2(maxLeft[i], maxRight[i]) - height[i]
		if buff > 0 {
			sum += buff
		}
	}

	return sum
}

/*柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。求在该柱状图中，能够勾勒出来的矩形的最大面积。

输入：heights = [2,1,5,6,2,3]
输出：10

记录每个柱子左边右边第一个小于该柱子的下标
*/

func largestRectangleArea(heights []int) int {
	maxLeft, minRight, result, length := make([]int, len(heights)), make([]int, len(heights)), 0, len(heights)

	maxLeft[0], minRight[length-1] = -1, length
	for i := 1; i < length; i++ {
		bar := i - 1
		for bar >= 0 && heights[bar] >= heights[i] {
			bar = maxLeft[bar] // if this bar is higher than i : get the first smaller bar of this bar
		}
		maxLeft[i] = bar
	}

	for i := length - 2; i >= 0; i-- {
		bar := i + 1
		for bar <= length-1 && heights[bar] >= heights[i] {
			bar = minRight[bar]
		}
		minRight[i] = bar
	}

	for i := 0; i < length; i++ {
		sum := heights[i] * (minRight[i] - maxLeft[i] - 1)
		result = MaxOf2(sum, result)
		fmt.Println(minRight[i], maxLeft[i], result)
	}

	return result
}
