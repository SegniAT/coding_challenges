#include <bits/stdc++.h>
#include <vector>

using namespace std;

int main() {
  int n, h;
  cin >> n >> h;
  vector<int> arr;

  while (n > 0) {
    int el;
    cin >> el;
    arr.push_back(el);
    n--;
  }

  int maxWidth = 0;

  for (auto el : arr) {
    if (el > h) {
      maxWidth += 2;
    } else {
      maxWidth++;
    }
  }

  cout << maxWidth + n << "\n";
  return 0;
}
