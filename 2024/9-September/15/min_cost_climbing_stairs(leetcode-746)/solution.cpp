#include <bits/stdc++.h>
#include <climits>

using namespace std;

class Solution {
public:
  int topDown(int n, vector<int> &cost, vector<int> &dp) {
    if (dp[n] != INT_MAX)
      return dp[n];
    if (n == 0 || n == 1)
      return 0;
    int res = min(topDown(n - 1, cost, dp) + cost[n - 1],
                  topDown(n - 2, cost, dp) + cost[n - 2]);
    dp[n] = res;
    return res;
  }

  int bottomUp(vector<int> &cost) {
    int N = cost.size();
    vector<int> dp(N + 1, INT_MAX);
    dp[0] = 0;
    dp[1] = 0;

    for (int i = 2; i < N; i++) {
      dp[i] = min(dp[i - 1] + cost[i - 1], dp[i - 2] + cost[i - 2]);
    }

    return dp[N];
  }

  int minCostClimbingStairs(vector<int> &cost) {
    // int N = cost.size();
    //  cost to reach
    // vector<int> dp(N + 2, INT_MAX);
    // return topDown(N + 1, cost, dp);
    return bottomUp(cost);
  }
};

int main() {
  auto s = new Solution();
  vector<int> inp = {10, 15, 20};

  cout << "ans: " << s->minCostClimbingStairs(inp);
}
