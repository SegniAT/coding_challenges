#include <bits/stdc++.h>

using namespace std;
class Solution {
public:
  int uniquePaths(int m, int n) {
    vector<vector<int>> memo(m + 1, vector<int>(n + 1));
    memo[1][1] = 1;
    for (int r = 1; r < m + 1; r++) {
      for (int c = 1; c < n + 1; c++) {
        if (r == 1 && c == 1)
          continue;
        memo[r][c] = memo[r - 1][c] + memo[r][c - 1];
      }
    }

    return memo[m][n];
  }
};

int main() {}
