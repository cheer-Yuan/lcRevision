package main

import (
	"fmt"
	"math"
	"sort"
)

//// time : average O(nlogn), worst O(n2)
//func QSParint(a []int, min, max int)  {
//
//}
//
//func QSint(a []int) []int {
//
//
//}

func Partition(list []int, low, high int) int {
	// may use a random method to set the pivot
	pivot := list[low]

	for low < high {
		//pivot < pointerx
		for low < high && pivot <= list[high] {
			high--
		}
		list[low] = list[high]

		//pivot > pointer
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
		fmt.Println(low, high, list)
	}

	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		pivot := Partition(list, low, high)

		QuickSort(list, low, pivot-1)
		QuickSort(list, pivot+1, high)
	}
}

/*最小绝对差
给你个整数数组 arr，其中每个元素都 不相同。
请你找到所有具有最小绝对差的元素对，并且按升序的顺序返回。
*/
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	results := [][]int{}
	minDiff := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] < minDiff {
			results = [][]int{}
			minDiff = arr[i] - arr[i-1]
			results = append(results, []int{arr[i-1], arr[i]})
		} else if minDiff == arr[i]-arr[i-1] {
			results = append(results, []int{arr[i-1], arr[i]})
		}
	}

	return results
}
