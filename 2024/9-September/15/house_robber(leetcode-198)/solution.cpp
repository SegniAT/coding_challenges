#include <bits/stdc++.h>
#include <climits>
#include <vector>

using namespace std;

class Solution {
public:
  int topDown(int h, vector<int> &nums, vector<int> &dp) {
    if (h < 0) {
      return 0;
    }
    if (dp[h] != INT_MAX)
      return dp[h];
    int opt1 = topDown(h - 2, nums, dp) + nums[h];
    int opt2 = topDown(h - 3, nums, dp) + nums[h];
    dp[h] = max(opt1, opt2);
    return dp[h];
  }
  int bottomUP(vector<int> &nums) {
    int N = nums.size();

    vector<int> dp(N, INT_MAX);
    for (int i = 0; i < N; i++) {
      int opt1 = 0, opt2 = 0;
      if (i - 2 >= 0) {
        opt1 = dp[i - 2];
      }
      if (i - 3 >= 0) {
        opt2 = dp[i - 3];
      }

      dp[i] = max(opt1, opt2) + nums[i];
    }

    for (auto a : dp) {
      cout << a << ", ";
    }

    if (N == 1) {
      return dp[N - 1];
    }

    return max(dp[N - 1], dp[N - 2]);
  }

  int rob(vector<int> &nums) {
    // int N = nums.size();
    // vector<int> dp(N, INT_MAX);
    // return max(topDown(N - 1, nums, dp), topDown(N - 2, nums, dp));
    return bottomUP(nums);
  }
};
int main() {
  vector<int> nums = {1, 2, 3, 1};
  auto s = new Solution();
  cout << s->rob(nums) << "\n";
  nums = {1, 20};
  cout << s->rob(nums) << "\n";
}
