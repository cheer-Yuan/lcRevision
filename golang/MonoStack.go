package main

/*
给定一个整数数组temperatures，表示每天的温度，返回一个数组answer，其中answer[i]是指在第 i 天之后，才会有更高的温度。如果气温在这之后都不会升高，请在该位置用0 来代替。
输入: temperatures = [73,74,75,71,69,72,76,73]
输出: [1,1,4,2,1,1,0,0]

通常是一维数组，要寻找任一个元素的右边或者左边第一个比自己大或者小的元素的位置，此时我们就要想到可以用单调栈了。
单调栈的本质是空间换时间，因为在遍历的过程中需要用一个栈来记录右边第一个比当前元素？的元素，优点是只需要遍历一次。
*/

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	monostack := StackOfInt{}

	for i := 0; i < len(temperatures); i++ {
		if i == 0 {
			monostack.push(i)
			continue
		}
		for temperatures[i] > temperatures[monostack.peek()] { // if bigger than the first in the stack : pop and renew the answer
			pop := monostack.pop()
			result[pop] = i - pop // calculate the distance
			if monostack.size() == 0 {
				break
			}
		}
		monostack.push(i)
	}
	return result
}

/*
nums1中数字x的 下一个更大元素 是指x在nums2 中对应位置 右侧 的 第一个 比x大的元素。
给你两个 没有重复元素 的数组nums1 和nums2 ，下标从 0 开始计数，其中nums1是nums2的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。

输入：nums1 = [4,1,2], nums2 = [1,3,4,2].
输出：[-1,3,-1]
解释：nums1 中每个值的下一个更大元素如下所述：
- 4 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。
- 1 ，用加粗斜体标识，nums2 = [1,3,4,2]。下一个更大元素是 3 。
- 2 ，用加粗斜体标识，nums2 = [1,3,4,2]。不存在下一个更大元素，所以答案是 -1 。

*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, len(nums1))
	monostack := StackOfInt{}

	// no redundant elements => use map to store and quest
	dic := map[int]int{}
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
		dic[nums1[i]] = i
	}

	for i := 1; i < len(nums2); i++ {
		for nums2[i] > nums2[monostack.peek()] {
			index, ifexist := dic[nums2[monostack.peek()]]
			if ifexist == true {
				result[index] = nums2[i]
			}
			monostack.pop()
		}

		monostack.push(i)

	}

}