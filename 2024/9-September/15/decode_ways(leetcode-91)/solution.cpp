#include <bits/stdc++.h>

using namespace std;

class Solution {
public:
  int topDown(string s, int ind, vector<int> &memo) {
    int N = s.size();
    if (ind == N) {
      return 1;
    }

    if (memo[ind] != INT_MAX)
      return memo[ind];

    char currChar = s[ind];

    int ways = 0;
    if ('1' <= currChar && currChar <= '9') {
      ways += topDown(s, ind + 1, memo);
    }

    if (ind + 1 < N &&
        ((s[ind] == '1' && '0' <= s[ind + 1] && s[ind + 1] <= '9') ||
         (s[ind] == '2' && '0' <= s[ind + 1] && s[ind + 1] <= '6'))) {
      ways += topDown(s, ind + 2, memo);
    }
    memo[ind] = ways;
    return ways;
  }

  int bottomUp(string s) {
    int N = s.size();
    vector<int> memo(N);

    for (int i = 0; i < N; i++) {
      char currChar = s[i];
      // option 1
      if ('1' <= currChar && currChar <= '9') {
        if (i - 1 < 0) {
          memo[i] = 1;
        } else {
          memo[i] += memo[i - 1];
        }
      }

      // option 2
      if (i - 1 >= 0 &&
          ((s[i - 1] == '1' && '0' <= currChar && currChar <= '9') ||
           (s[i - 1] == '2' && '0' <= currChar && currChar <= '6'))) {
        if (i - 2 < 0) {
          memo[i] += 1;
        } else {
          memo[i] += memo[i - 2];
        }
      }
    }
    return memo[N - 1];
  }

  int numDecodings(string s) {
    // int N=s.size();
    // vector<int> memo    (N, INT_MAX);
    // return topDown(s, 0, memo);
    return bottomUp(s);
  }
};
int main() {}
