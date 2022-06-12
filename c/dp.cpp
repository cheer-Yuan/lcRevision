//
// Created by cheer on 2022/6/11.
//

#include <vector>
#include <string>
using namespace std;

/*将字符串翻转到单调递增
如果一个二进制字符串，是以一些 0（可能没有 0）后面跟着一些 1（也可能没有 1）的形式组成的，那么该字符串是 单调递增 的。
给你一个二进制字符串 s，你可以将任何 0 翻转为 1 或者将 1 翻转为 0 。
返回使 s 单调递增的最小翻转次数。

第i位的答案与第i-1位相关，联想使用动态规划。
分析：第i位可以是0或1
dp[i][0] = dp[i-1][0] + if(s[i] == 1) : 第i位若是0，则第i-1位必须是0。
dp[i][1] = min(dp[i-1][1], dp[i-1][0]) + if(s[i] == 0) : 若第i位为1，则第i-1位可以是0或1。
*/
class Solution {
public:
    int minFlipsMonoIncr(string s) {
        vector<int> dp = {0, 0};
        for (int i = 0; i < s.size(); ++i) {
            int dp0, dp1 = 0;
            if (s[i] == '0') {
                dp0 = dp[0];
                dp1 = min(dp[0], dp[1]) + 1;
            } else {
                dp0 = dp[0] + 1;
                dp1 = min(dp[0], dp[1]);
            }
            dp = {dp0, dp1};
        }
        return min(dp[0], dp[1]);
    }
};
