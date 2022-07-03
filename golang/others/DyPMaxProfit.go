package others

import "fmt"

/*
给定一个数组 prices ，它的第i 个元素prices[i] 表示一支给定股票第 i 天的价格。

你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

示例 1：

输入：[7,1,5,3,6,4]
输出：5
解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
示例 2：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这种情况下, 没有交易完成, 所以最大利润为 0。

提示：

1 <= prices.length <= 105
0 <= prices[i] <= 104

dp[i]：在第i天能获得的最高利润，
min ： 最低价
if price[i] - min  > dp[i] ：更新
初始化：0
*/

func maxProfit1(prices []int) int {
	OldProfit, NewProfit := 0, 0

	min := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		}

		NewProfit = MaxOf2(OldProfit, prices[i] - min)
		OldProfit = NewProfit
	}


	return NewProfit
}

/*
给定一个数组 prices ，其中prices[i] 是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

*/

func maxProfit2(prices []int) int {
	Profit, min := 0, prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < prices[i - 1] {
			Profit += prices[i - 1] - min

		} else {
			Profit += prices[i] - prices[i - 1]
		}
		min = prices[i]
	}

	return Profit
}


/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成两笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

示例1:

输入：prices = [3,3,5,0,0,3,1,4]
输出：6
解释：在第 4 天（股票价格 = 0）的时候买入，在第 6 天（股票价格 = 3）的时候卖出，这笔交易所能获得利润 = 3-0 = 3 。
    随后，在第 7 天（股票价格 = 1）的时候买入，在第 8 天 （股票价格 = 4）的时候卖出，这笔交易所能获得利润 = 4-1 = 3 。
示例 2：

输入：prices = [1,2,3,4,5]
输出：4
解释：在第 1 天（股票价格 = 1）的时候买入，在第 5 天 （股票价格 = 5）的时候卖出, 这笔交易所能获得利润 = 5-1 = 4 。
    注意你不能在第 1 天和第 2 天接连购买股票，之后再将它们卖出。
    因为这样属于同时参与了多笔交易，你必须在再次购买前出售掉之前的股票。
示例 3：

输入：prices = [7,6,4,3,1]
输出：0
解释：在这个情况下, 没有交易完成, 所以最大利润为 0。
示例 4：

输入：prices = [1]
输出：0

478 143 579

提示：
1 <=prices.length <= 105
0 <=prices[i] <=105

思路：
考虑四个变量：对于第i天：
第一次买入可以获得的最大的收益 j = 1
	dp[i][1] = 通过两种方式到达：买入：dp[i - 1][0] - price[i]；持有：dp[i - 1][1]，max选择
第一次卖出可以获得的最大的收益 j = 2
	dp[i][2] = 卖出: dp[i - 1][1] + price[i]；之前卖出了：dp[i - 1][2]
第二次买入可以获得的最大的收益 j = 3
	dp[i][3] = 买入：dp[i - 1][2] - price[i]；持有：dp[i - 1][3]
第二次卖出可以获得的最大的收益 j = 4
	dp[i][4] = 卖出：dp[i - 1][3] + price[i]；持有:dp[i - 1][4]

初始化：dp[0][0] = 0, dp[0][1] dp[0][1] = dp[0][3] = - price[0]
*/

func maxProfit3(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 5)
	}

	dp[0][1], dp[0][3] = - prices[0], - prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i - 1][0]
		dp[i][1] = MaxOf2(dp[i - 1][0] - prices[i], dp[i - 1][1])
		dp[i][2] = MaxOf2(dp[i - 1][1] + prices[i], dp[i - 1][2])
		dp[i][3] = MaxOf2(dp[i - 1][2] - prices[i], dp[i - 1][3])
		dp[i][4] = MaxOf2(dp[i - 1][3] + prices[i], dp[i - 1][4])
	}

	return dp[len(prices) - 1][4]
}


/*
给定一个整数数组prices ，它的第 i 个元素prices[i] 是一支给定的股票在第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。



示例 1：

输入：k = 2, prices = [2,4,1]
输出：2
解释：在第 1 天 (股票价格 = 2) 的时候买入，在第 2 天 (股票价格 = 4) 的时候卖出，这笔交易所能获得利润 = 4-2 = 2 。
示例 2：

输入：k = 2, prices = [3,2,6,5,0,3]
输出：7
解释：在第 2 天 (股票价格 = 2) 的时候买入，在第 3 天 (股票价格 = 6) 的时候卖出, 这笔交易所能获得利润 = 6-2 = 4 。随后，在第 5 天 (股票价格 = 0) 的时候买入，在第 6 天 (股票价格 = 3) 的时候卖出, 这笔交易所能获得利润 = 3-0 = 3 。


提示：

0 <= k <= 100
0 <= prices.length <= 1000
0 <= prices[i] <= 1000
*/

func maxProfit4(k int, prices []int) int {
	if k == 0 || len(prices) == 0 {
		return 0
	}

	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2 * k + 1)
	}

	for i := 1; i < k * 2; i += 2 {
		dp[0][i] = - prices[0]
	}

	fmt.Println(dp[0][1], dp[0][3])

	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i - 1][0]
		for j := 1; j <= k * 2; j += 2 {
			dp[i][j] = MaxOf2(dp[i - 1][j - 1] - prices[i], dp[i - 1][j])
			dp[i][j + 1] = MaxOf2(dp[i - 1][j] + prices[i], dp[i - 1][j + 1])
		}
	}

	return dp[len(prices) - 1][k * 2]
}


/*
给定一个整数数组，其中第i个元素代表了i的股票价格 。

设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:

你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。

思路：对于day i :
	状态 j = 0：买入
		操作1：买入后一直持有 dp[i - 1][0]
		操作2：当天买入：
			前一天是冷东期：dp[i - 1][3] - prices[i]
			前一天非冷冻期：dp[i - 1][1] - prices[i]
	状态 j = 1：卖出后冷冻期已过
		操作1：一直是冷冻期已过：dp[i - 1][1]
		操作2：前一天是冷冻期：dp[i - 1][3]
	状态 j = 2：当天卖出
		操作：只能由j = 0转化而来：dp[i - 1][0] + prices[i]
	状态 j = 3：冷冻期
		操作：只能由j = 2转化而来：dp[i - 1][2]
初始化：j = 0 : - prices[i]
	j = 1 : 0
	j = 2 : 0
	j = 3 : 0
*/


func maxProfit5(prices []int) int {
	leng := len(prices)
	dp := make([][4]int, leng)
	dp[0][0] = - prices[0]

	for i := 1; i < leng; i++ {
		dp[i][0] = MaxOf2(dp[i - 1][0], MaxOf2(dp[i - 1][3] - prices[i], dp[i - 1][1] - prices[i]))
		dp[i][1] = MaxOf2(dp[i - 1][1], dp[i - 1][3])
		dp[i][2] = dp[i - 1][0] + prices[i]
		dp[i][3] = dp[i - 1][2]
	}

	return MaxOf2(dp[leng - 1][0], MaxOf2(dp[leng - 1][1], MaxOf2(dp[leng - 1][2], dp[leng - 1][3])))
}


/*
给定一个整数数组prices，其中第i个元素代表了第i天的股票价格 ；整数fee 代表了交易股票的手续费用。

你可以无限次地完成交易，但是你每笔交易都需要付手续费。如果你已经购买了一个股票，在卖出它之前你就不能再继续购买股票了。

返回获得利润的最大值。

思路：对于day i ：
	状态j = 0： 买入
		操作1：持有 dp[i - 1][0]
		操作2：买入 dp[i - 1][1] - prices[i] - i
	状态j = 1： 卖出
		操作1：已经卖出：dp[i - 1][1]
		操作2：当天卖出：dp[i - i][0] + prices[i]
*/

func maxProfit(prices []int, fee int) int {
	leng := len(prices)
	dp := make([][2]int, leng)
	dp[0][0] = - prices[0] - fee

	for i := 1; i < leng; i++ {
		dp[i][0] = MaxOf2(dp[i - 1][0], dp[i - 1][1] - prices[i] - fee)
		dp[i][1] = MaxOf2(dp[i - 1][1], dp[i - 1][0] + prices[i])
	}

	return MaxOf2(dp[leng - 1][0], dp[leng - 1][1])
}

