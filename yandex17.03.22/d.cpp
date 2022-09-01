#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <algorithm>

using namespace std;

int main() {
	int n;
	cin >> n;
	string str;
	cin >> str;

	int min_bal = 1;
	int max_bal = -1;
	int min_bal_ind = -1;
	int max_bal_ind = -1;

	int balance = 0;
	map<char, int> val_bind = {{'0', -1}, {'1', 1}};
	for (int i = 0; i < n; i++) {
		// balance management
		int val = val_bind[str[i]];
		balance += val;

		int res = -1;
		// calc
		if (i > 0)
			if (str[i] == '0') {
				if (balance < 0) res = 1;
				else if (balance - max_bal < 0) res = max_bal_ind + 1;
			}
			else {
				if (balance > 0) res = 1;
				else if (balance - min_bal > 0) res = min_bal_ind + 1;
			}
		cout << res << " ";

		balance -= val;
		// index updating
		if (balance < min_bal) {
			min_bal = balance;
			min_bal_ind = i;
		}
		if (balance > max_bal) {
			max_bal = balance;
			max_bal_ind = i;
		}
		balance += val;
	}
}