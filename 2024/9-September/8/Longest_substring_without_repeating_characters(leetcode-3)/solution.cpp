#include <bits/stdc++.h>
using namespace std;

class Solution {
public:
  int lengthOfLongestSubstring(string s) {
    int N = s.length();
    map<char, int> cache;
    int left = 0, right = 0, longest = 0;

    while (right < N) {
      cache[s[right]]++;
      while (cache[s[right]] > 1) {
        cache[s[left]]--;
        left++;
      }
      longest = max(longest, right - left + 1);
      right++;
    }

    return longest;
  }
};

int main() {};
