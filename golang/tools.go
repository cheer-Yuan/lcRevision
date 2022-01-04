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
		return - i
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

	return float64(partInt) + float64(parsDec) / 100
}

