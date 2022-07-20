package main

/*1260. 二维网格迁移
给你一个 m 行 n列的二维网格grid和一个整数k。你需要将grid迁移k次。
每次「迁移」操作将会引发下述活动：
位于 grid[i][j]的元素将会移动到grid[i][j + 1]。
位于grid[i][n- 1] 的元素将会移动到grid[i + 1][0]。
位于 grid[m- 1][n - 1]的元素将会移动到grid[0][0]。
请你返回k 次迁移操作后最终得到的 二维网格。

cad：所有列往右移，右下元素往左上，最右列往左下
按行一维展开后观察，一次移位相当于往右移动一次
*/
func shiftGrid(grid [][]int, k int) [][]int {
	rowCount := len(grid)
	columnCount := len(grid[0])
	totalCount := rowCount * columnCount

	result := make([][]int, rowCount)
	for resultRows := 0; resultRows < rowCount; resultRows++ { // map内存
		result[resultRows] = make([]int, columnCount)
	}

	shiftCount := k % totalCount
	for oldIndex := 0; oldIndex < totalCount; oldIndex++ { // 按照一维数据位置处理
		newIndex := oldIndex + shiftCount
		newIndex %= totalCount
		oldRow, newRow := handleIndex(oldIndex/columnCount, newIndex/columnCount, rowCount)
		oldColumn, newColumn := handleIndex(oldIndex%columnCount, newIndex%columnCount, columnCount)
		result[newRow][newColumn] = grid[oldRow][oldColumn]
	}

	return result
}

func handleIndex(oldRow, newRow, row int) (int, int) {
	if oldRow == row {
		oldRow = 0
	}
	if newRow == row {
		newRow = 0
	}

	return oldRow, newRow
}
