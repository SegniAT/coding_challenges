#include <bits/stdc++.h>

typedef long long ll;

using namespace std;

int solve(int n, int H, vector<ll> arr) {
  ll largest[2] = {0, 0};
  for (auto num : arr) {
    if (num >= largest[1]) {
      ll temp = largest[1];
      largest[0] = temp;
      largest[1] = num;
    } else if (num >= largest[0]) {
      largest[0] = num;
    }
  }
  ll x = largest[0], y = largest[1];

  // NOTE: the 2x is added b/c we're using two numbers/two steps
  if (H % (x + y) == 0) {
    return 2 * (H / (x + y));
  } else if (H % (x + y) <= y) {
    return 2 * (H / (x + y)) + 1;
  } else {
    return 2 * (H / (x + y)) + 2;
  }

  return 0;
}

int main() {
  int t;
  scanf("%d", &t);

  for (; t > 0; t--) {
    int n, H;
    vector<ll> arr;
    scanf("%d %d", &n, &H);

    for (int i = 0; i < n; i++) {
      ll x;
      scanf("%lld", &x);
      arr.push_back(x);
    }
    printf("%d\n", solve(n, H, arr));
  }
}
