package main

import (
	"fmt"
	"strings"
)

func ReverseString(s string) string {
	a := []rune(s)

	for i, j := 0, len(a) - 1; i < j; i, j = i + 1, j - 1 {
		a[i], a[j] = a[j], a[i]
	}
	return string(a)
}

func SubPanlidrome(s string) string {
	RevS := ReverseString(s)
	Len := len(s)

	var Sub [][]int

	for i := 0; i < Len; i++ {
		a := make([]int, Len)
		Sub = append(Sub, a)
	}

	for i := 0; i < Len; i++ {
		if s[i] == RevS[0] {
			Sub[i][0] = 1
		}
		if s[0] == RevS[i] {
			Sub[0][i] = 1
		}
	}

	fmt.Println(Sub)

	maxL, maxI := 0, 0
	for i := 1; i < Len; i++ {
		for j := 1; j < Len; j++ {
			if s[i] == RevS[j] {
				Sub[i][j] = Sub[i-1][j-1] + 1
				if Sub[i][j] > maxL {
					if Len - j  -1 + Sub[i][j] - 1  == i {
						maxL = Sub[i][j]
						maxI = i
					}
				}
			}
		}
	}


	var Result []string
	for i := 0; i < maxL; i++ {
		Result = append(Result, string(s[maxI - maxL + 1 + i]))
	}

	if maxL == 0 {
		return string(s[0])
	}

	return strings.Join(Result, "")
}