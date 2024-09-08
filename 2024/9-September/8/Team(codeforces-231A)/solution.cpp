#include "bits/stdc++.h"

using namespace std;

int main() {
  int n, sln = 0;
  cin >> n;

  while (n > 0) {
    int sure = 0;
    int temp;
    for (int i = 0; i < 3; ++i) {
      cin >> temp;
      sure += temp;
    }
    if (sure >= 2) {
      sln++;
    }
    n--;
  }
  cout << sln << "\n";
}
