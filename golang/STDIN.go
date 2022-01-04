package main

import (
	"bufio"
	"fmt"
	"os"
)

func STDIN(s string) int {
	// 读取所有stdin到一个字符串
	r := bufio.NewReader(os.Stdin)
	input, _ := r.ReadBytes('\n')
	fmt.Println(input)

	// 读取input到变量
	var a int
	fmt.Scan(&a)


	return 0

}

