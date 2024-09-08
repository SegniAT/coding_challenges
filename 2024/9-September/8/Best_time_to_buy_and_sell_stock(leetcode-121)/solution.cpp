#include <algorithm>
#include <bits/stdc++.h>

using namespace std;

int main() {}

// first soln
/*
class Solution {
public:
  int maxProfit(vector<int> &prices) {
    int N = prices.size();
    int prevMin = 1e5;
    int maxProf = 0;

    for (int i = 0; i < N; i++) {
      maxProf = max(maxProf, prices.at(i) - prevMin);
      prevMin = min(prevMin, prices.at(i));
    }

    return maxProf;
  }
};
         */
// left just tracks the current minimum
class Solution {
public:
  int maxProfit(vector<int> &prices) {
    int N = prices.size();
    int maxProf = 0, left = 0, right = 1;
    while (right < N) {
      if (prices.at(left) < prices.at(right)) {
        maxProf = max(maxProf, prices.at(right) - prices.at(left));
      } else {
        left = right;
      }
      right++;
    }

    return maxProf;
  }
};
