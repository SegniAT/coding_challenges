#include <bits/stdc++.h>
#include <vector>

using namespace std;

class Solution {
public:
  int climbStairsBottomUp(int n) {
    if (n == 1 || n == 2)
      return n;

    int dp[n];
    dp[0] = 1;
    dp[1] = 2;

    for (int i = 3; i < n; i++) {
      dp[i] = dp[i - 1] + dp[i - 2];
    }
    return dp[n - 1];
  }

  int climbStairsTopDown(int n, vector<int> &dp) {
    if (dp[n] != 0) {
      return dp[n];
    }
    if (n == 1 || n == 2)
      return n;
    int res = climbStairsTopDown(n - 1, dp) + climbStairsTopDown(n - 2, dp);
    dp[n] = res;
    return res;
  }

  int climbStairs(int n) {
    vector<int> dp(n + 1, 0);
    return climbStairsTopDown(n, dp);
  }
};
