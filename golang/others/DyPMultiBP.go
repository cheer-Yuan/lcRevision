package others

/*
背包最大重量为10。

物品为：

重量	价值	数量
物品0	1	15	2
物品1	3	20	3
物品2	4	30	2

问可装的最大价值
*/

func multiBP() int {
	// limits here
	weight := []int{1, 3, 4}
	value := []int{15, 20, 30}
	nums := []int{2, 3, 2}
	m := 10

	// solution
	dp := make([]int, m + 1)
	dp[0] = 1

	for index, i := range weight {
		for j := m; j >= i; i-- {
			for k := 1; k < nums[index] && j - k * weight[index] >= 0; k++ {
				dp[j] = MaxOf2(dp[j], dp[j - i] + k * value[index] )
			}
		}
	}

	return dp[m]
}
