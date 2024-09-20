#include <bits/stdc++.h>
#include <unordered_set>

using namespace std;

int diversity(vector<int> notes) {
  unordered_set<int> set;
  for (int note : notes)
    set.insert(note);
  return set.size();
}

// wrong
int topDown(vector<int> notes, int ind, vector<int> &memo) {
  if (ind < 0) {
    return diversity(notes);
  }

  if (memo[ind] != INT_MAX)
    return memo[ind];

  int res = 0;
  res = max(res, topDown(notes, ind - 1, memo));
  notes[ind]++;
  res = max(res, topDown(notes, ind - 1, memo));

  memo[ind] = res;

  return res;
}

int greedy(vector<int> notes) {
  unordered_set<int> set;
  for (auto note : notes) {
    if (set.count(note))
      note++;
    set.insert(note);
  }

  return set.size();
}

int main() {
  ios::sync_with_stdio(0);
  cin.tie(0);

  int t;
  cin >> t;

  while (t-- > 0) {
    int n;
    cin >> n;

    vector<int> notes(n);
    vector<int> memo(n, INT_MAX);
    for (int i = 0; i < n; i++) {
      cin >> notes[i];
    }

    // int res = topDown(notes, n - 1, memo);
    int res = greedy(notes);
    cout << res << "\n";
  }
}
