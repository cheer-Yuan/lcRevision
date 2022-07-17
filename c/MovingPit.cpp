//
// Created by zhiyu on 2022/7/16.
//

/*
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。

实现 MovingAverage 类：
MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val)成员函数 next每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。
*/
#include <queue>
using namespace std;


class MovingAverage {
public:
    int Sum;
    int Size;
    queue<int> numList;

    MovingAverage(int size) {
        Sum = 0;
        Size = size;
        numList = queue<int>();
    }

    double sum(int val) {
        if(numList.size() < Size) {
            numList.push(val);
        } else {
            Sum -= numList.front();
            numList.push(val);
            numList.pop();
        }
        Sum += val;

        return double(numList.size());
    }

    double next(int val) {
        double numDiv = sum(val);
        return double(Sum) / numDiv;
    }
};
