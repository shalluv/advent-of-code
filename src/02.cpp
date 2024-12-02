#include <bits/stdc++.h>
using namespace std;

int main() {
  ifstream f("./input/02.txt");
  if (!f.is_open()) {
    return 1;
  }

  string s;
  int cnt = 0;
  while (getline(f, s)) {
    istringstream iss(s);
    int prev = -1, curr;
    int dir = 0;
    bool succ = true;
    string c;
    while (iss >> c) {
      if (!succ) continue;
      curr = stoi(c);
      if (prev == -1) {
        prev = curr;
        continue;
      }
      if (prev == curr || abs(prev - curr) > 3) {
        succ = false;
        continue;
      }
      if (dir == 0) {
        dir = (curr > prev) ? 1 : -1;
      } else if ((dir == -1 && curr > prev) || (dir == 1 && curr < prev)) {
        succ = false;
      }
      prev = curr;
    }
    if (succ) ++cnt;
  }

  cout << cnt << '\n';

  return 0;
}