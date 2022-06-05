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
class Solution {
private:
    mt19937 gen{random_device{}()};     // included in random
    // 用法：
    // std::mt19937 mt_rand(std::random_device{}());
    // std::mt19937 mt_rand(time(0));
    // std::mt19937 mt_rand(std::chrono::system_clock::now().time_since_epoch().count());
    uniform_real_distribution<double> dis;  // 默认返回 double 型浮点值的连续分布。可以按如下方式生成一个返回值在范围 [_, _) 内的分布对象：
    double xc, yc, r;

public:
    Solution(double radius, double x_center, double y_center): dis(-radius, radius), xc(x_center), yc(y_center), r(radius) {}   // constructor

    vector<double> randPoint() {
        while (true) {
            double x = dis(gen), y = dis(gen);
            if (x * x + y * y < r * r) {
                return {x + xc, y + yc};
            }
        }
    }
};