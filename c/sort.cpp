//
// Created by cheer on 2022/6/13.
//

#include <vector>
#include <stdlib.h>
#include <string>
#include <unordered_map>
using namespace std;

/*高度检查器
校打算为全体学生拍一张年度纪念照。根据要求，学生需要按照 非递减 的高度顺序排成一行。
排序后的高度情况用整数数组 expected 表示，其中 expected[i] 是预计排在这一行中第 i 位的学生的高度（下标从 0 开始）。
给你一个整数数组 heights ，表示 当前学生站位 的高度情况。heights[i] 是这一行中第 i 位学生的高度（下标从 0 开始）。
返回满足 heights[i] != expected[i] 的 下标数量 。
*/
class Solution {
public:
    void QuickSort(vector<int>& nums, int left, int right) {
        if (left < right) {
            int i = left, j = right, pivot = nums[left];
            while (i < j) {
                while (i < j && pivot <= nums[j]) j--;
                if (i < j) nums[i++] = nums[j];

                while (i < j  && pivot > nums[i]) i++;
                if (i < j) nums[j--] = nums[i];
            }
            nums[i] = pivot;

            QuickSort(nums, left, i - 1);
            QuickSort(nums, i + 1, right);
        }
    }

    int heightChecker(vector<int>& heights) {
        int result = 0;
        vector<int> sorted = {};
        for (int i = 0; i < heights.size(); ++i) {
            sorted.emplace_back(heights[i]);
        }
        QuickSort(sorted, 0, sorted.size() - 1);

        for (int i = 0; i < sorted.size(); ++i) {
            if (heights[i] != sorted[i]) result++;
        }

        return result;
    }
};