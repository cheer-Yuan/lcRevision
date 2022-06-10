//
// Created by zhiyu on 2022/6/7.
//

#include <vector>
using namespace std;

/*爱吃香蕉的珂珂
珂珂喜欢吃香蕉。这里有 n 堆香蕉，第 i 堆中有piles[i]根香蕉。警卫已经离开了，将在 h 小时后回来。
珂珂可以决定她吃香蕉的速度 k （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 k 根。如果这堆香蕉少于 k 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。
返回她可以在 h 小时内吃掉所有香蕉的最小速度 k（k 为整数）。

思路：如果珂珂在 hh 小时内吃掉所有香蕉的最小速度是每小时 k 个香蕉，则当吃香蕉的速度大于每小时 k 个香蕉时一定可以在 h 小时内吃掉所有香蕉，当吃香蕉的速度小于每小时 k 个香蕉时一定不能在 h 小时内吃掉所有香蕉。存在绝对关系，考虑二分法。
*/
class minEatingSpeedSolution {
public:
    static int minEatingSpeed(vector<int>& piles, int h) {
        int low = 1;
        int high = 0;
        for (int pile : piles) {
            high = max(high, pile);         // 取速度上界：最多香蕉量
        }

        int k = high;
        while (low < high) {                        // 当速度下界小于上界
            int speed = (high - low) / 2 + low;     // 二分
            long time = getTime(piles, speed);
            if (time <= h) {
                k = speed;
                high = speed;                       // 有余量 ： 设为上界
            } else {
                low = speed + 1;                    // 无余量： 下界+1
            }
        }
        return k;
    }

    static long getTime(const vector<int>& piles, int speed) {      // 吃掉所有香蕉的时间
        long time = 0;
        for (int pile : piles) {
            int curTime = (pile + speed - 1) / speed;       // // 等价于 ceiling( pile / speed)，吃掉此堆香蕉的时间
            time += curTime;
        }
        return time;
    }
};


/* 搜索旋转排序数组
整数数组 nums 按升序排列，数组中的值 互不相同 。
在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为[4,5,6,7,0,1,2] 。
给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。

直接遍历的时间复杂度高
*/
class searchSolution {
public:
    static int search(vector<int>& nums, int target) {
        if (nums.empty()) return -1;
        if (nums.size() == 1) return nums[0] == target ? 0 : -1;

        int low = 0;
        int high = nums.size() - 1;

        while (low <= high) {                    // 寻找旋转点
            int med = (low + high) / 2;
            if (nums[med] == target) return med;
            if (nums[0] <= nums[med]) {        // 旋转点位于高区
                if (target < nums[med] && target >= nums[0]) {
                    high = med - 1;     // 目标位于低区
                } else {
                    low = med + 1;
                }
            }
            else {                              // 旋转点位于低区
                if (target > nums[med] && target <= nums[nums.size() - 1]) {
                    low = med + 1;                // 目标位于高区
                } else {
                    high = med - 1;
                }
            }
        }

        return -1;
    }
};


/*
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
如果数组中不存在目标值 target，返回[-1, -1]。
复杂度为O(log n)

输入：nums = [5,7,7,8,8,10], target = 8
输出：[3,4]

输入：nums = [], target = 0
输出：[-1,-1]
*/
class searchRangeSolution {
public:
    static int binarySearch(vector<int>& nums, int target, bool ifRight) {
        int low = 0;
        int high = nums.size() - 1;

        while (low <= high) {
            int med = (low + high) / 2;
            if (ifRight) {
                if (nums[med] == target && (med == nums.size() - 1 || nums[med + 1] != target)) return med;
                if (nums[med] > target) high = med - 1;
                if (nums[med] <= target) low = med + 1;
            } else {
                if (nums[med] == target && (med == 0 || nums [med - 1] != target)) return med;
                if (nums[med] < target) low = med + 1;
                if (nums[med] >= target) high = med - 1;
            }
        }

        return -1;
    }

    static vector<int> searchRange(vector<int>& nums, int target) {
        if (nums.empty()) return vector<int>{-1, -1};
        if (nums.size() == 1) return nums[0] == target ? vector<int>{0, 0} : vector<int>{-1, -1};

        return vector<int>{binarySearch(nums, target, false), binarySearch(nums, target, true)};
    }
};