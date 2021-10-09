package main
//
//// time : average O(nlogn), worst O(n2)
//func QSParint(a []int, min, max int)  {
//
//}
//func QSint(a []int) []int {
//
//
//}

func Partition(list []int, low, high int) int {
	// may use a random method to set the pivot
	pivot := list[low]

	for low < high {
		//pivot < pointer
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
	}

	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low{
		pivot := Partition(list, low, high)

		QuickSort(list, low, pivot - 1)
		QuickSort(list, pivot + 1, high)
	}
}