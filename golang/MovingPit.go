package main

/*
给定一个整数数据流和一个窗口大小，根据该滑动窗口的大小，计算滑动窗口里所有数字的平均值。

实现 MovingAverage 类：
MovingAverage(int size) 用窗口大小 size 初始化对象。
double next(int val)成员函数 next每次调用的时候都会往滑动窗口增加一个整数，请计算并返回数据流中最后 size 个值的移动平均值，即滑动窗口里所有数字的平均值。
*/

type MovingAverage struct {
	Sum     int
	Size    int
	NumList []int
}

/** Initialize your data structure here. */
func MovingAverageConstructor(size int) MovingAverage {
	return MovingAverage{
		Sum:     0,
		Size:    size,
		NumList: make([]int, 0),
	}
}

func (this *MovingAverage) push(val int) {
	this.NumList = append(this.NumList[1:], val)
}

func (this *MovingAverage) sum(val int) int {
	if len(this.NumList) < this.Size {
		this.Sum += val
		this.NumList = append(this.NumList, val)
	} else {
		this.Sum -= this.NumList[0]
		this.push(val)
		this.Sum += val
	}

	return len(this.NumList)
}

func (this *MovingAverage) Next(val int) float64 {
	numDiv := this.sum(val)
	return float64(this.Sum) / float64(numDiv)
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
