//
// Created by zhiyu on 2022/6/1.
//

#include <vector>
using namespace std;

#include <numeric>     // accumulate 函数
/*
accumulate函数进行自定义求和：
struct Grade {
    string Name;
    int note;
};
int main() {
    Grade subjects[3] = {
            {"Eng", 10},
            {"Mat", 20},
            {"Phy", 30}
    };

    int sum = accumulate(subjects, subjects + 3, 0, [](int a, const Grade& b){return a + b.note;});
    cout << sum << endl;
}
*/

#include <algorithm>   // sort 函数
#include <string>
#include <iostream>


/*火柴拼正方形
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i个火柴棒的长度。你要用 所有的火柴棍拼成一个正方形。你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
如果你能使这个正方形，则返回 true ，否则返回 false 。
*/
class Solutionmakesquare {
public:
    bool dfs(vector<int> &edges, vector<int> &matchsticks, int totalLen, int iMatch) {
        if (iMatch == matchsticks.size()) {
            return true;
        }

        for (int i = 0; i < edges.size(); ++i) {
            edges[i] += matchsticks[iMatch];
            if (edges[i] <= totalLen && dfs(edges, matchsticks, totalLen, iMatch + 1)) return true;
            edges[i] -= matchsticks[iMatch];
        }

        return false;
    }

    bool makesquare(vector<int> &matchsticks) {
        int totalLen = accumulate(matchsticks.begin(), matchsticks.end(), 0);
        if (totalLen % 4 != 0) return false;
        sort(matchsticks.begin(), matchsticks.end(), greater<>());       // 从大到小排序

        vector<int> edges(4);
        return dfs(edges, matchsticks, totalLen / 4, 0);
    }
};

