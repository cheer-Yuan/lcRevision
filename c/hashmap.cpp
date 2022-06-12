//
// Created by cheer on 2022/6/12.
//

#include <vector>
#include <stdlib.h>
#include <string>
#include <unordered_map>
using namespace std;

/*
你有一个单词列表words和一个模式pattern，你想知道 words 中的哪些单词与模式匹配。
如果存在字母的排列 p，使得将模式中的每个字母 x 替换为 p(x) 之后，我们就得到了所需的单词，那么单词与模式是匹配的。
（回想一下，字母的排列是从字母到字母的双射：每个字母映射到另一个字母，没有两个字母映射到同一个字母。）
返回 words 中与给定模式匹配的单词列表。
你可以按任何顺序返回答案。

输入：words = ["abc","deq","mee","aqq","dkd","ccc"], pattern = "abb"
输出：["mee","aqq"]
解释：
"mee" 与模式匹配，因为存在排列 {a -> m, b -> e, ...}。
"ccc" 与模式不匹配，因为 {a -> c, b -> c, ...} 不是排列。
因为 a 和 b 映射到同一个字母。
*/
class Solution {
public:
    vector<string> findAndReplacePattern(vector<string>& words, string pattern) {
        vector<string> result = {};
        for (auto &word : words) {
            bool ifOK = true;
            unordered_map<char, char> map1, map2 = {};
            for (int i = 0; i < pattern.size(); ++i) {
                if (!map1.count(pattern[i]) || !map2.count(word[i])) {      // count方法：检查该元素是否存在于map中
                    if (!map1.count(pattern[i])) map1[pattern[i]] = word[i];
                    if (!map2.count(word[i])) map2[word[i]] = pattern[i];
                }

                if (map1[pattern[i]] != word[i] || map2[word[i]] != pattern[i]) {
                    ifOK = false;
                    break;
                }
            }
            if (ifOK) result.emplace_back(word);
        }

        return result;
    }
};
