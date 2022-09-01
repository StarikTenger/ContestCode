#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <map>
#include <set>

using namespace std;


int main() {
	int n;
	cin >> n;

	int min_up = 1000;
	int max_down = 0;
	int sum = 0;

	for (int i = 0; i < n; i++) {
		int val;
		cin >> val;

		if (i % 2 == 0) {
			sum += val;
			min_up = min(min_up, val);
		}
		else {
			sum -= val;
			max_down = max(max_down, val);
		}
	}

	int sum_new = sum;
	sum_new += max_down * 2 - min_up * 2;
	cout << max(sum, sum_new);
}