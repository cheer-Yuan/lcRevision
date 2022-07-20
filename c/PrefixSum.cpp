//
// Created by cheer on 2022/7/20.
//

#include <vector>
using namespace std;

/*303. 区域和检索 - 数组不可变
给定一个整数数组 nums，处理以下类型的多个查询:

计算索引left和right（包含 left 和 right）之间的 nums 元素的 和 ，其中left <= right
实现 NumArray 类：

NumArray(int[] nums) 使用数组 nums 初始化对象
int sumRange(int i, int j) 返回数组 nums中索引left和right之间的元素的 总和 ，包含left和right两点（也就是nums[left] + nums[left + 1] + ... + nums[right])
*/
class NumArray {
public:
    vector<int> sums;

    NumArray(vector<int>& nums) {
        sums.resize(nums.size() + 1);
        int sum = 0;
        for (int indexSums = 1; indexSums < sums.size(); ++indexSums) {
            sum += nums[indexSums - 1];
            sums[indexSums] = sum;
        }
    }

    int sumRange(int left, int right) {
        return sums[right + 1] - sums[left];
    }
};