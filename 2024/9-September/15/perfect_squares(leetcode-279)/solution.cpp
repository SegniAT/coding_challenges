#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int helper(int n, vector<int> &dp) {
    if (n == 0) {
      return 0;
    }

    if (dp[n] != INT_MAX) {
      return dp[n];
    }

    int m = INT_MAX;
    for (int i = 1; i * i <= n; i++) {
      m = min(m, helper(n - i * i, dp) + 1);
    }
    dp[n] = m;
    return dp[n];
  }

  int numSquares(int n) {
    vector<int> dp(n + 1, INT_MAX);
    return helper(n, dp);
  }
};
