#include <bits/stdc++.h>

using namespace std;

int main() {
  int n;
  string games;
  cin >> n >> games;

  int anton = 0, danik = 0;
  for (auto game : games) {
    if (game == 'A') {
      anton++;
    } else {
      danik++;
    }
  }

  if (anton > danik) {
    cout << "Anton\n";
  } else if (anton < danik) {
    cout << "Danik\n";
  } else {
    cout << "Friendship\n";
  }
}
