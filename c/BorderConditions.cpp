//
// Created by cheer on 2022/6/14.
//

#include <vector>
#include <iostream>
using namespace std;


/*对角线遍历
给你一个大小为 m x n 的矩阵 mat ，请以对角线遍历的顺序，用一个数组返回这个矩阵中的所有元素。

可优化：遍历方向由坐标之和决定
*/
class Solution {
public:
    vector<int> findDiagonalOrder(vector<vector<int>>& mat) {
        int nRow = mat.size(), nCol = mat[0].size();
        vector<int> result = {};

        if (nRow == 1) return mat[0];
        if (nCol == 1) {
            for (auto i : mat) result.emplace_back(i[0]);
            return result;
        }

        int i = 0, j = 0, step = 1;
        bool ifReverse = true;
        while (i < nRow && j < nCol) {
            result.emplace_back(mat[i][j]);
            if (i == nRow - 1 && j == nCol - 1)  return result;
            if (ifReverse && (((i == 0 || i == nRow - 1)) || (j == 0 || j == nCol - 1))) {
                ifReverse = false;
                step *= -1;
                if ((i == 0 || i == nRow - 1) && j < nCol - 1) {
                    j++;
                    continue;
                }
                if ((j == 0 || j == nCol - 1) && i < nRow - 1) {
                    i++;
                    continue;
                }
            }

            ifReverse = true;
            i -= step;
            j += step;
        }

        return result;
    }
};