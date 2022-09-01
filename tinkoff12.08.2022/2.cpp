#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <map>

using namespace std;

int main() {
	int n;
	cin >> n;

	map<string, int> team_score;
	int max_score = 0;

	for (int i = 0; i < n; i++) {
		vector<string> names(3);

		cin >> names[0] >> names[1] >> names[2];
		sort(names.begin(), names.end());
		string team = names[0] + names[1] + names[2];

		if (!team_score.count(team))
			team_score[team] = 0;
		team_score[team]++;

		max_score = max(team_score[team], max_score);
	}

	cout << max_score;
}