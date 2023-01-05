#include <iostream>
#include <map>
#include <string>
#include <vector>

// https://leetcode.cn/problems/24-game/

using namespace std;

class Solution {
public:
    std::string getKey(const std::vector<int> &arr) {
        std::string out;
        for (const auto &tmp : arr) {
            out += "#" + std::to_string(tmp);
        }
        return out;
    }

    // 全排列，输出到inputVector
    void permu(std::vector<int> &input, int index, std::map<std::string, std::vector<int>> &inputVector) {
        if (index == 3) {
            inputVector[getKey(input)] = input;
            return;
        }
        permu(input, index + 1, inputVector);
        for (int i = index + 1; i < 4; i++) {
            int tmp = input[index];
            input[index] = input[i];
            input[i] = tmp;
            permu(input, index + 1, inputVector);
            input[i] = input[index];
            input[index] = tmp;
        }
    }

    // 返回加减乘除的结果，除法存在余数就不计算在内
    std::vector<double> getValue(double a, int b) {
        std::vector<double> out;
        out.emplace_back(a + b);
        out.emplace_back(a - b);
        out.emplace_back(a * b);
        out.emplace_back(a * 1.0 / b);
        return out;
    }

    // 计算一个input的所有结果是否为24
    bool compute(const std::vector<int> &input) {
        std::vector<double> lastResult;  // 上一位计算结果
        lastResult.emplace_back(input[0]);
        for (size_t i = 1; i < 4; i++) {
            std::vector<double> tmpResult;
            for (const auto &tmp : lastResult) {
                auto tmpV = getValue(tmp, input[i]);
                if (i == 3) {
                    // 做一下提前退出
                    for (const auto &item : tmpV) {
                        if (item - 24 < 1e-6 && item - 24 > -1e-6) {
                            // 打印一下顺序
                            for (const auto &item1 : input) {
                                cout << item1 << " ";
                            }
                            cout << endl;
                            return true;
                        }
                    }
                } else {
                    tmpResult.insert(tmpResult.end(), tmpV.begin(), tmpV.end());
                }
            }
            lastResult = std::move(tmpResult);
        }
        return false;
    }

    bool judgePoint24(vector<int> &cards) {
        std::map<std::string, std::vector<int>> inputVector;  // 全排列数组
        permu(cards, 0, inputVector);

        for (const auto &tmp : inputVector) {
            if (compute(tmp.second)) {
                return true;
            }
        }
        return false;
    }
};
