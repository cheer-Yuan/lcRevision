//
// Created by zhiyu on 2022/6/4.
//

#include <vector>
#include <cmath>
#include <random>
using namespace std;




/*在圆内随机生成点
给定圆的半径和圆心的位置，实现函数 randPoint ，在圆中产生均匀随机点。

实现Solution类:
Solution(double radius, double x_center, double y_center)用圆的半径radius和圆心的位置 (x_center, y_center) 初始化对象
randPoint()返回圆内的一个随机点。圆周上的一点被认为在圆内。答案作为数组返回 [x, y] 。
*/
/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
class randPointSolution {
private:
    mt19937 gen{random_device{}()};     // included in random
    // 用法：
    // std::mt19937 mt_rand(std::random_device{}());
    // std::mt19937 mt_rand(time(0));
    // std::mt19937 mt_rand(std::chrono::system_clock::now().time_since_epoch().count());
    uniform_real_distribution<double> dis;  // 默认返回 double 型浮点值的连续分布。可以按如下方式生成一个返回值在范围 [_, _) 内的分布对象：
    double xc, yc, r;

public:
    randPointSolution(double radius, double x_center, double y_center): dis(-radius, radius), xc(x_center), yc(y_center), r(radius) {}   // constructor

    vector<double> randPoint() {
        while (true) {
            double x = dis(gen), y = dis(gen);
            if (x * x + y * y < r * r) {
                return {x + xc, y + yc};
            }
        }
    }
};


/*非重叠矩形中的随机点
给定一个由非重叠的轴对齐矩形的数组 rects ，其中 rects[i] = [ai, bi, xi, yi] 表示 (ai, bi) 是第 i 个矩形的左下角点，(xi, yi) 是第 i 个矩形的右上角角点。设计一个算法来随机挑选一个被某一矩形覆盖的整数点。矩形周长上的点也算做是被矩形覆盖。所有满足要求的点必须等概率被返回。
在一个给定的矩形覆盖的空间内任何整数点都有可能被返回。
请注意，整数点是具有整数坐标的点。

实现Solution类:
Solution(int[][] rects)用给定的矩形数组rects 初始化对象。
int[] pick() 返回一个随机的整数点 [u, v] 在给定的矩形所覆盖的空间内。
*/
class Solution {
private:
    mt19937 gen{random_device{}()};
    uniform_int_distribution<int> dis;      // 生成整数[a, b]
    vector<vector<int>>& rects;
    vector<int> sums;

public:
    Solution(vector<vector<int>>& rects) : rects{rects}, dis(0, sums.back() - 1) {
        this->sums.emplace_back(0);     // add a value to the vector
        for (auto &rect : rects) {
            this->sums.emplace_back(sums.back() + (rect[2] - rect[0] + 1) * (rect[3] - rect[1] + 1));    // back : take the last value
        }                 // 循环体中修改a，b中对应内容也会修改
    }

    static int binarySearch(vector<int>& nums, int target) {
        int low = 0;
        int high = nums.size() - 1;

        while (low < high) {
            int mid = (low + high / 2);
            if (mid < nums.size() - 1 && nums[mid] <= target && nums[mid + 1] > target || mid == nums.size() - 1) return mid;
            if (nums[mid] > target) high = mid - 1;
            if (nums[mid] < target) low = mid + 1;
        }
        return -1;
    }

    vector<int> pick() {
        int points = dis(gen);
        int rectIndex = binarySearch(sums, points);
        int modP = points - sums[rectIndex];
        int X = modP % (this->rects[rectIndex][2] - this->rects[rectIndex][0] + 1);
        int Y = abs(modP / (this->rects[rectIndex][2] - this->rects[rectIndex][0] + 1));
        return vector<int>{this->rects[rectIndex][0] + X, this->rects[rectIndex][1] + Y};
    }
};
/**
 * Your Solution object will be instantiated and called as such:
 * Solution* obj = new Solution(rects);
 * vector<int> param_1 = obj->pick();
 */