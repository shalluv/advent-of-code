#include <bits/stdc++.h>
using namespace std;

int main() {
  ifstream f("./input/01.txt");
  if (!f.is_open()) {
    return 1;
  }

  string s;
  vector<int> v1(1000), v2(1000);
  int i = 0;
  while (getline(f, s)) {
    v1[i] = stoi(s.substr(0, s.find(' ')));
    v2[i] = stoi(s.substr(s.find(' ') + 1));
    ++i;
  }
  sort(v1.begin(), v1.end());
  sort(v2.begin(), v2.end());

  long long dist = 0;
  for (int i = 0; i < 1000; ++i) {
    dist += abs(v1[i] - v2[i]);
  }

  cout << dist << '\n';

  return 0;
}