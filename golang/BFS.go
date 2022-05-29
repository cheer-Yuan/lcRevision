package main

import "sort"

/*
BFS使用队列，把每个还没有搜索到的点依次放入队列，然后再弹出队列的头部元素当做当前遍历点。

BFS总共有两个模板：

模板一：
如果不需要确定当前遍历到了哪一层，只需要访问完所有节点就可以时。
while queue 不空：
    cur = queue.pop()
    if cur 有效且未被访问过：
        进行处理
    for 节点 in cur 的所有相邻节点：
        if 该节点有效：
            queue.push(该节点)

模板二：
如果要确定当前遍历到了哪一层，需要知道最少移动步数时，BFS 模板如下。
这里增加了 level 表示当前遍历到二叉树中的哪一层了，也可以理解为在一个图中，现在已经走了多少步了。size 表示在当前遍历层有多少个元素，也就是队列中的元素数，我们把这些元素一次性遍历完，即把当前层的所有元素都向外走了一步。
level = 0
while queue 不空：
    size = queue.size()
    while (size --) {
        cur = queue.pop()
        if cur 有效且未被访问过：
            进行处理
        for 节点 in cur的所有相邻节点：
            if 该节点有效：
                queue.push(该节点)
    }
    level ++;
*/

/*为高尔夫比赛砍树
你被请来给一个要举办高尔夫比赛的树林砍树。树林由一个m x n 的矩阵表示， 在这个矩阵中：
0 表示障碍，无法触碰
1表示地面，可以行走
比 1 大的数表示有树的单元格，可以行走，数值表示树的高度
每一步，你都可以向上、下、左、右四个方向之一移动一个单位，如果你站的地方有一棵树，那么你可以决定是否要砍倒它。你需要按照树的高度从低向高砍掉所有的树，每砍过一颗树，该单元格的值变为 1（即变为地面）。你将从 (0, 0) 点开始工作，返回你砍完所有树需要走的最小步数。 如果你无法砍完所有的树，返回 -1 。可以保证的是，没有两棵树的高度是相同的，并且你至少需要砍倒一棵树。

输入：forest = [[2,3,4],[0,0,5],[8,7,6]]
输出：6
解释：可以按与示例 1 相同的路径来砍掉所有的树。
(0,0) 位置的树，可以直接砍去，不用算步数。

思路：
根据题意，总的移动步数为：从起点到最低的树的最少步数 + 从最低的树到第 2 低的树的最少步数 + 从第 2 低的树到第 3 低的树的最少步数 + ... 直至所有树被砍完。

首先对矩阵中的树按照树的高度进行排序，我们依次求出相邻的树之间的最短距离。
利用广度优先搜索，按照层次遍历，处理队列中的节点（网格位置）。
记录在某个时间点已经添加到队列中的节点，这些节点已被处理或在等待处理的队列中。对于下一个要处理的每个节点，查看他们的四个方向上相邻的点，如果相邻的点没有被遍历过且不是障碍，将其加入到队列中，直到找到终点为止，返回当前的步数即可。最终返回所有的步数之和即为最终结果。
*/
func cutOffTree(forest [][]int) int {
	dir4 := []struct{x, y int}{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	type pair struct {
		dis, x, y int
	}
	trees := []pair{}														// 树的列表
	for iRow, Row := range forest{
		for iCol, Col := range Row {
			if Col > 1 {
				trees = append(trees, pair{Col, iRow, iCol})		// 如果一个格子有树，加入列表
			}
		}
	}
	sort.Slice(trees, func(i, j int) bool {
		return trees[i].dis < trees[j].dis
	})																		// 按照树的高度进行排序

	bfs := func(sx, sy, tx, ty int) int {
		m, n := len(forest), len(forest[0])
		vis := make([][]bool, m)
		for i := range vis {
			vis[i] = make([]bool, n)
		}
		vis[sx][sy] = true
		q := []pair{{0, sx, sy}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			if p.x == tx && p.y == ty {
				return p.dis
			}
			for _, d := range dir4 {
				if x, y := p.x+d.x, p.y+d.y; 0 <= x && x < m && 0 <= y && y < n && !vis[x][y] && forest[x][y] > 0 {
					vis[x][y] = true
					q = append(q, pair{p.dis + 1, x, y})
				}
			}
		}
		return -1
	}

	preX, preY, ans := 0, 0, 0
	for _, t := range trees {
		d := bfs(preX, preY, t.x, t.y)
		if d < 0 {
			return -1
		}
		ans += d
		preX, preY = t.x, t.y
	}
	return ans
}