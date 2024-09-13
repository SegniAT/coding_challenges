#include <bits/stdc++.h>

using namespace std;

int main() {
  // ios::sync_with_stdio(0);
  // cin.tie(0);

  int t;
  scanf("%d", &t);

  while (t-- > 0) {
    int n;
    scanf("%d", &n);
    getchar(); // Consume the leftover newline from scanf

    int balloons = 0;
    unordered_set<char> set;
    string s;
    getline(cin, s);
    for (char c : s) {
      if (set.count(c) == 0) {
        balloons += 2;
        set.insert(c);
      } else {
        balloons += 1;
      }
    }
    printf("%d\n", balloons);
  }

  return 0;
}
