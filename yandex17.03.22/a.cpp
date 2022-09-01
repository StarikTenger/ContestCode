#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <algorithm>

using namespace std;

int main() {
	int n;
	cin >> n;
	vector<int> nums(n);
	for (auto& x : nums) cin >> x;
	
	string alphabet = "abcdefghijklmnopqrstuvwxyz ";
	int prev = 0;
	for (int i = 0; i < n; i++) {
		int diff = prev ^ nums[i];
		int ind = 0;
		while (diff = diff>>1, diff != 0) ind++;
		cout << alphabet[ind];
		prev = nums[i];
	}
}