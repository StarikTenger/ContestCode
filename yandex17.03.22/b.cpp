#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <algorithm>

using namespace std;

bool is_letter(char s) {
	return s >= 'A' && s <= 'Z';
}

int main() {
	int n, m;
	cin >> n >> m;
	vector<string> field(n);
	for (auto& s : field) cin >> s;
	
	vector<string> field_new = field;
	for (auto& s : field_new) for (int j = 0; j < m; j++) s[j] = '.';
	
	for (int i = 0; i < n; i++) {
		for (int j = 0; j < m; j++) {
			int i1 = n - i - 1;
			int j1 = m - j - 1;
			if (is_letter(field[i][j])) {
				field_new[i1][j1] = field[i][j];
				field_new[i1 - 1][j1] = '_';
				field_new[i1 + 1][j1] = '_';
				field_new[i1][j1 - 1] = '/';
				field_new[i1][j1 + 1] = '\\';
				field_new[i1 + 1][j1 - 1] = '\\';
				field_new[i1 + 1][j1 + 1] = '/';
			}
		}
	}

	for (auto& s : field_new) cout << s << "\n";
}