package main

import (
	"strconv"
	"strings"
)

func MaxOf2(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}

func MinOf2(i, j int) int {
	if i > j {
		return j
	} else {
		return i
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	} else {
		return i
	}
}

func swap(s string, i, j int) string {
	if i > j {
		return s[0:j] + string(s[i]) + s[j+1:i] + string(s[j]) + s[i+1:len(s)]
	} else if i < j {
		return s[0:i] + string(s[j]) + s[i+1:j] + string(s[i]) + s[j+1:len(s)]
	} else {
		return s
	}
}

// 四舍五入保留两位
func trunc(f float64) float64 {
	n := strconv.FormatFloat(f, 'f', -1, 64)
	newn := strings.Split(n, ".")

	partInt, _ := strconv.Atoi(newn[0])
	parsDec, _ := strconv.Atoi(newn[1][0:2])

	if newn[1][2] >= 53 {
		parsDec++
	}

	return float64(partInt) + float64(parsDec)/100
}

// 矩阵幂乘
func matPower(n, size int, mat [][]int) [][]int {
	result := [][]int{} // 初始化
	for i := 0; i < size; i++ {
		temp := []int{}
		for j := 0; j < size; j++ {
			temp = append(temp, mat[i][j])
		}
		result = append(result, mat[i])
	}

	iter := 0
	for iter < n-1 {
		power := [][]int{}
		for row := 0; row < size; row++ {
			prow := []int{}
			for col := 0; col < size; col++ {
				temp := 0
				for i := 0; i < size; i++ {
					temp += result[row][i] * mat[i][col]
				}
				prow = append(prow, temp)
			}
			power = append(power, prow)
		}
		result = power
		iter++
	}
	return result
}

func sub2D(a, b []int) []int {
	return []int{a[0] - b[0], a[1] - b[1]}
}

// 计算向量的叉积 pq x qr， 如果叉积小于 0，可以知道向量顺时针旋转， 否则逆时针旋转。
func cross(p, q, r []int) int {
	return (q[0]-p[0])*(r[1]-q[1]) - (q[1]-p[1])*(r[0]-q[0])
}

func distance2D(p, q []int) int {
	return (p[0]-q[0])*(p[0]-q[0]) + (p[1]-q[1])*(p[1]-q[1])
}
