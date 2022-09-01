#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
	int left = 100;
	int right = -100;
	int up = 100;
	int down = -100;

	for (int i = 0; i < 4; i++) {
		int x, y;
		cin >> x >> y;
		left = min(left, x);
		right = max(right, x);
		up = min(up, y);
		down = max(down, y);
	}

	int side = max(right - left, up - down);
	cout << side * side;
}