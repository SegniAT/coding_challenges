#include <algorithm>
#include <bits/stdc++.h>
#include <climits>
#include <vector>

using namespace std;

class Solution {
public:
  int greedy(string colors, vector<int> &neededTime) {
    int res = 0;
    int rSum = 0, cMax = 0;
    char cChar = ' ';
    for (int i = 0; i < int(colors.size()); i++) {
      char c = colors[i];
      if (c != cChar) {
        res += rSum - cMax;
        rSum = 0;
        cMax = 0;
      }
      cChar = c;
      rSum += neededTime[i];
      cMax = max(cMax, neededTime[i]);
    }
    res += rSum - cMax;
    return res;
  }

  bool isColorful(string colors) {
    int N = colors.size();
    if (N == 1)
      return true;

    for (int i = 1; i < N; i++) {
      if (colors[i] == colors[i - 1])
        return false;
    }
    return true;
  }

  int topDown(string colors, vector<int> &neededTime, map<string, int> &dp) {

    int N = colors.size();

    if (dp.find(colors) != dp.end()) {
      return dp[colors];
    }
    if (isColorful(colors)) {
      return 0;
    }

    int res = INT_MAX;

    for (int i = 0; i < N; i++) {
      string nColors = colors.substr(0, i) + colors.substr(i + 1);
      int rem = topDown(nColors, neededTime, dp) + neededTime[i];

      res = min(res, rem);
    }

    dp[colors] = res;
    return res;
  }

  int minCost(string colors, vector<int> &neededTime) {
    // return greedy(colors, neededTime);
    map<string, int> dp;
    int res = topDown(colors, neededTime, dp);
    for (auto e : dp) {
      cout << e.first << ": " << e.second << "\n";
    }
    return res;
  }
};

int main() {
  string colors = "bbbaaa";
  vector<int> neededTime = {4, 9, 3, 8, 8, 9};
  auto s = new Solution();
  int sol = s->minCost(colors, neededTime);
  cout << "sol: " << sol << "\n";
}
