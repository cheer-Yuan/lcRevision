//
// Created by cheer on 2022/7/20.
//

#include <vector>
using namespace std;

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

class Solution {
public:
    vector<vector<int>> shiftGrid(vector<vector<int>>& grid, int k) {
        int rowCount = grid.size();
        int colCount = grid[0].size();
        int allCount = rowCount * colCount;

        vector<vector<int>> result(rowCount, vector<int>(colCount));

        int shiftCount = k % allCount;
        for (int oldIndex = 0; oldIndex < allCount; ++oldIndex) {
            int newIndex = (oldIndex + shiftCount) % allCount;
            int oldRow = handelIndex(oldIndex / colCount, rowCount);
            int newRow = handelIndex(newIndex / colCount, rowCount);
            result[newRow][newIndex % colCount] = grid[oldRow][oldIndex % colCount];
        }

        return result;
    }

    int handelIndex(int oldIndex, int limit) {
        if (oldIndex >= limit) oldIndex = 0;
        return oldIndex;
    }
};