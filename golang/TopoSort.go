package main

import "fmt"

/*课程表
你这个学期必须选修 numCourses 门课程，记为0到numCourses - 1 。
在选修某些课程之前需要一些先修课程。 先修课程按数组prerequisites 给出，其中prerequisites[i] = [ai, bi] ，表示如果要学习课程ai 则 必须 先学习课程 bi 。
例如，先修课程对[0, 1] 表示：想要学习课程 0 ，你需要先完成课程 1 。
请你判断是否可能完成所有课程的学习？如果可以，返回 true ；否则，返回 false 。

示例 2：
输入：numCourses = 2, prerequisites = [[1,0],[0,1]]
输出：false
解释：总共有 2 门课程。学习课程 1 之前，你需要先完成课程 0 ；并且学习课程 0 之前，你还应先完成课程 1 。这是不可能的。


 拓扑排序：
我们将每一门课看成一个节点；如果想要学习课程 A 之前必须完成课程 B，那么我们从 B 到 A 连接一条有向边。这样以来，在拓扑排序中，B 一定出现在 A 的前面。
DFS ：对图进行一遍深度优先搜索。当每个节点进行回溯的时候，我们把该节点放入栈中。最终从栈顶到栈底的序列就是一种拓扑排序。
*/
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses) 								// 制作有向邻接表
	for _, pairs := range prerequisites {
		graph[pairs[1]] = append(graph[pairs[1]], pairs[0])					// 加入所有边
	}

	var dfs func(node int)
	visited, valid := make([]int, numCourses), true
	dfs = func(node int) {									// dfs 遍历图，检测是否存在环
		visited[node] = 1									// 标记为已遍历
		for _, linkedNode := range graph[node] {
			if visited[linkedNode] == 0 {					// 尚未遍历到
				dfs(linkedNode)									// dfs递归
			} else if visited[linkedNode] == 1 {
				valid = false									// 第二次遍历到某个节点 == 检测到环
				return
			}
		}
		visited[node] = 2									// 已经完全遍历
	}

	for i := 0; i < numCourses && valid; i++ {				// 对于每个尚未遍历到的节点进行dfs
		if visited[i] == 0 {
			dfs(i)
		}
	}

	return valid
}
/*课程表 II
现在你总共有 numCourses 门课需要选，记为0到numCourses - 1。给你一个数组prerequisites ，其中 prerequisites[i] = [ai, bi] ，表示在选修课程 ai 前 必须 先选修bi 。
返回你为了学完所有课程所安排的学习顺序。可能会有多个正确的顺序，你只要返回 任意一种 就可以了。如果不可能完成所有课程，返回 一个空数组 。
*/
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	for _, i := range prerequisites {
		graph[i[1]] = append(graph[i[1]], i[0])
	}

	result := []int{}
	visited, valid := make([]int, numCourses), true
	var dfs func(node int)
	dfs = func(node int) {
		fmt.Println(node)
		visited[node] = 1
		for _, linkedNode := range graph[node] {
			if visited[linkedNode] == 0 {
				dfs(linkedNode)
			} else if visited[linkedNode] == 1 {
				fmt.Println(node)
				valid = false
				return
			}
		}
		visited[node] = 2
		result = append(result, node)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}

	if !valid {
		fmt.Println(1)
		return []int{}
	} else {
		for i := 0 ; i < numCourses / 2; i++ {
			result[i], result[numCourses - 1 - i] = result[numCourses - 1 - i], result[i]
		}
		return result
	}
}


/*猫和老鼠
两位玩家分别扮演猫和老鼠，在一张 无向 图上进行游戏，两人轮流行动。
图的形式是：graph[a] ，由邻接点数组标示
老鼠从节点 1 开始，第一个出发；猫从节点 2 开始，第二个出发。在节点 0 处有一个洞。
在每个玩家的行动中，他们 必须 沿着图中与所在当前位置连通的一条边移动。例如，如果老鼠在节点 1 ，那么它必须移动到 graph[1] 中的任一节点。
此外，猫无法移动到洞中（节点 0）。
然后，游戏在出现以下三种情形之一时结束：
如果猫和老鼠出现在同一个节点，猫获胜。
如果老鼠到达洞中，老鼠获胜。
如果某一位置重复出现（即，玩家的位置和移动顺序都与上一次行动相同），游戏平局。

给你一张图 graph ，并假设两位玩家都都以最佳状态参与游戏：
如果老鼠获胜，则返回1；
如果猫获胜，则返回 2；
如果平局，则返回 0


拓扑排序：消除结果和轮数之间的关系，从边界情况出发遍历其他情况。根据当前的移动方，可以得到上一轮的所有可能状态：
	如果当前的移动方是老鼠，则上一轮的移动方是猫，上一轮状态中老鼠所在节点是 mouse，猫所在节点可能是 graph[cat] 中的任意一个节点（除了节点 0）；
	如果当前的移动方是猫，则上一轮的移动方是老鼠，上一轮状态中老鼠所在节点可能是 graph[mouse] 中的任意一个节点，猫所在节点是 cat。
所有状态的结果都初始化为平局，   对于上一轮的移动方，只有当可以确定上一轮状态是必胜状态或者必败状态时，才更新上一轮状态的结果。
	如果上一轮的移动方和当前状态的获胜方相同，由于当前状态为上一轮的移动方的必胜状态，因此上一轮的移动方一定可以移动到当前状态而获胜，上一轮状态为上一轮的移动方的必胜状态。
	如果上一轮的移动方和当前状态的获胜方不同，则上一轮的移动方需要尝试其他可能的移动，可能有以下三种情况
		如果存在一种移动可以到达上一轮的移动方的必胜状态，则上一轮状态为上一轮的移动方的必胜状态；
		如果所有的移动都到达上一轮的移动方的必败状态，则上一轮状态为上一轮的移动方的必败状态；
		如果所有的移动都不能到达上一轮的移动方的必胜状态，但是存在一种移动可以到达上一轮的移动方的未知状态，则上一轮状态为上一轮的移动方的未知状态。
遍历过程中，从当前状态出发遍历上一轮的所有可能状态，如果上一轮状态的结果是平局且上一轮的移动方和当前状态的结果的获胜方不同，则将上一轮状态的度减 1。如果上一轮状态的度减少到 0，则从上一轮状态出发到达的所有状态都是上一轮的移动方的必败状态，因此上一轮状态也是上一轮的移动方的必败状态。
在确定上一轮状态的结果（必胜或必败）之后，即可从上一轮状态出发，遍历其他结果是平局的状态。当没有更多的状态可以确定胜负结果时，遍历结束，此时即可得到初始状态的结果。

*/
func catMouseGame(graph [][]int) int {
	mouseTurn, catTurn := 0, 1			// mod 2 = 0 >>> mouse turn
	draw, mouseWin, catWin := 0, 1, 2

	n := len(graph)						// 节点数
	degrees := make([][][2]int, n)		// 度矩阵。表示状态：鼠位-猫位-轮次
	results := make([][][2]int, n)		// 状态矩阵

	for i := range degrees {			//初始化度矩阵
		degrees[i] = make([][2]int, n)
		results[i] = make([][2]int, n)
	}
	for i, to := range graph {
		for j := 1; j < n; j++ {
			degrees[i][j][mouseTurn] = len(to)			// 对于老鼠而言，初始的度为老鼠所在的节点的相邻节点数
			degrees[i][j][catTurn] = len(graph[j])
		}
	}
	for _, y := range graph[0] {
		for i := range degrees {
			degrees[i][y][catTurn]--					// 对于猫而言，初始的度为猫所在的节点的相邻且非节点 0 的节点数
		}
	}

	type state struct{ mouse, cat, turn int }	// 初始化状态矩阵
	q := []state{}								// 状态列表
	for j := 1; j < n; j++ {
		results[0][j][mouseTurn] = mouseWin		// 鼠位=0：鼠胜
		results[0][j][catTurn] = mouseWin
		q = append(q, state{0, j, mouseTurn}, state{0, j, catTurn})
	}
	for i := 1; i < n; i++ {
		results[i][i][mouseTurn] = catWin		// 鼠位=猫位，猫胜
		results[i][i][catTurn] = catWin
		q = append(q, state{i, i, mouseTurn}, state{i, i, catTurn})
	}

	getPrevStates := func(s state) (prevStates []state) {		// from bottom to top 根据当前状态计算前序状态
		if s.turn == mouseTurn {									// 当前状态鼠
			for _, prev := range graph[s.cat] {							// 对于猫位
				if prev != 0 {
					prevStates = append(prevStates, state{s.mouse, prev, catTurn})
				}
			}
		} else {													// 当前状态猫
			for _, prev := range graph[s.mouse] {
				prevStates = append(prevStates, state{prev, s.cat, mouseTurn})
			}
		}
		return
	}

	for len(q) > 0 {
		s := q[0]
		q = q[1:]
		result := results[s.mouse][s.cat][s.turn]
		for _, p := range getPrevStates(s) {
			prevMouse, prevCat, prevTurn := p.mouse, p.cat, p.turn
			if results[prevMouse][prevCat][prevTurn] == draw {
				canWin := result == mouseWin && prevTurn == mouseTurn || result == catWin && prevTurn == catTurn
				if canWin {
					results[prevMouse][prevCat][prevTurn] = result
					q = append(q, p)
				} else {
					degrees[prevMouse][prevCat][prevTurn]--
					if degrees[prevMouse][prevCat][prevTurn] == 0 {
						if prevTurn == mouseTurn {
							results[prevMouse][prevCat][prevTurn] = catWin
						} else {
							results[prevMouse][prevCat][prevTurn] = mouseWin
						}
						q = append(q, p)
					}
				}
			}
		}
	}
	return results[1][2][mouseTurn]

}


///*
//一只猫和一只老鼠在玩一个叫做猫和老鼠的游戏。
//它们所处的环境设定是一个rows x cols的方格 grid，其中每个格子可能是一堵墙、一块地板、一位玩家（猫或者老鼠）或者食物。
//
//玩家由字符'C'（代表猫）和'M'（代表老鼠）表示。
//地板由字符'.'表示，玩家可以通过这个格子。
//墙用字符'#'表示，玩家不能通过这个格子。
//食物用字符'F'表示，玩家可以通过这个格子。
//字符'C'，'M'和'F'在grid中都只会出现一次。
//猫和老鼠按照如下规则移动：
//
//老鼠 先移动，然后两名玩家轮流移动。
//每一次操作时，猫和老鼠可以跳到上下左右四个方向之一的格子，他们不能跳过墙也不能跳出grid。
//catJump 和mouseJump是猫和老鼠分别跳一次能到达的最远距离，它们也可以跳小于最大距离的长度。
//它们可以停留在原地。
//老鼠可以跳跃过猫的位置。
//游戏有 4 种方式会结束：
//
//如果猫跟老鼠处在相同的位置，那么猫获胜。
//如果猫先到达食物，那么猫获胜。
//如果老鼠先到达食物，那么老鼠获胜。
//如果老鼠不能在 1000 次操作以内到达食物，那么猫获胜。
//给你rows x cols的矩阵grid和两个整数catJump和mouseJump，双方都采取最优策略，如果老鼠获胜，那么请你返回true，否则返回 false。
//*/
//func canMouseWin(grid []string, catJump int, mouseJump int) bool {
//
//}