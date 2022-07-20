package main

// 场景；重复求不变数组的区间和

/*303. 区域和检索 - 数组不可变
给定一个整数数组 nums，处理以下类型的多个查询:

计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])
*/

type NumArray struct {
	nums []int
	sums []int
}

func NumArrayConstructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	sum := 0
	for indexNums := 1; indexNums < len(sums); indexNums++ {
		sum += nums[indexNums-1]
		sums[indexNums] = sum
	}

	return NumArray{
		nums: nums,
		sums: sums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	return this.sums[right+1] - this.sums[left]
}

/*
给定一个二维矩阵 matrix，以下类型的多个请求：

计算其子矩形范围内元素的总和，该子矩阵的 左上角 为 (row1,col1) ，右下角 为 (row2,col2) 。
实现 NumMatrix 类：

NumMatrix(int[][] matrix)给定整数矩阵 matrix 进行初始化
int sumRegion(int row1, int col1, int row2, int col2)返回 左上角 (row1,col1)、右下角(row2,col2) 所描述的子矩阵的元素 总和 。
*/

type NumMatrix struct {
	nums [][]int
	sums [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	sums := make([][]int, len(matrix)+1)
	sums[0] = make([]int, len(matrix[0])+1)
	for indexRow := 1; indexRow < len(matrix)+1; indexRow++ {
		sum := 0
		sums[indexRow] = make([]int, len(matrix[0])+1)
		for indexColumn := 1; indexColumn < len(matrix[0])+1; indexColumn++ {
			sum += matrix[indexRow-1][indexColumn-1]
			sums[indexRow][indexColumn] = sum
			if indexRow > 0 {
				sums[indexRow][indexColumn] += sums[indexRow-1][indexColumn]
			}
		}
	}

	return NumMatrix{
		nums: matrix,
		sums: sums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sums[row2+1][col2+1] - this.sums[row2+1][col1] - this.sums[row1][col2+1] + this.sums[row1][col1]
}

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的连续子数组的个数 。
*/
func subarraySum(nums []int, k int) int {

}
