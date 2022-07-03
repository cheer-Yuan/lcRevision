package others

import (
	"sort"
)

/*爱吃香蕉的珂珂
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

思路：如果珂珂在 hh 小时内吃掉所有香蕉的最小速度是每小时 k 个香蕉，则当吃香蕉的速度大于每小时 k 个香蕉时一定可以在 h 小时内吃掉所有香蕉，当吃香蕉的速度小于每小时 k 个香蕉时一定不能在 h 小时内吃掉所有香蕉。存在绝对关系，考虑二分法。
*/
func minEatingSpeed(piles []int, h int) int {
	maxspeed := 0
	for _, pile := range piles {		// 遍历确定上界
		if pile > maxspeed {
			maxspeed = pile
		}
	}

	return 1 + sort.Search(maxspeed - 1, func(speed int) bool {		// uses binary search to find and return the smallest index i in [0, n) at which f(i) is true
		speed++
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed	// 等价于 ceiling( pile / speed)
		}
		return time <= h
	})
}


/* 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。

直接遍历的时间复杂度高
*/
func search(nums []int, target int) int {
	if nums == nil {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	low, high := 0, len(nums) - 1

	for (low <= high) {
		med := (low + high) / 2
		if nums[med] == target {
			return med
		}
		if nums[0] <= nums[med] {
			if target >= nums[0] && target < nums[med] {
				high = med - 1
			} else {
				low = med + 1
			}
		} else {
			if target > nums[med] && target <= nums[len(nums) - 1] {
				low = med + 1
			} else {
				high = med - 1
			}
		}
	}
	return -1
}


/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
复杂度为O(log n)

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

输入：nums = [], target = 0
 输出：[-1,-1]
*/
func searchRange(nums []int, target int) []int {
	left := sort.SearchInts(nums, target)		// SearchInts searches for x in a sorted slice of ints.  如果不存在返回len或某个idnex
	if left == len(nums) || nums[left	] != target {
		return []int{-1, -1}
	}
	right := sort.SearchInts(nums, target + 1) - 1		// 一定会落在 min(nums[i] > n) or len(nums)
	return []int{left, right}
}

func isBoomerang(points [][]int) bool {
	a1, a2 := float64(points[1][1] - points[0][1]), float64(points[1][0] - points[0][0])
	b1, b2 := float64(points[2][1] - points[0][1]), float64(points[2][0] - points[0][0])
	if points[0][0] == points[1][0] && points[0][1] == points[1][1] || points[1][0] == points[2][0] && points[1][1] == points[2][1] || points[0][0] == points[2][0] && points[0][1] == points[2][1] || a1 / a2 == b1 / b2 {
		return false
	} else {
		return true
	}
}


