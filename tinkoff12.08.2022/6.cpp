#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <map>
#include <set>

using namespace std;

struct Floor {
	long long weight = 0;
	vector<long long> inputs;
};

int main() {
	long long n;
	cin >> n;

	map<long long, Floor> floors;

	for (long long i = 0; i < n; i++) {
		long long a, b;
		cin >> a >> b;

		if (!floors.count(a))
			floors[a] = {};
		if (!floors.count(b))
			floors[b] = {};

		if (a == b) {
			floors[a].weight++;
		}
		else {
			floors[b].inputs.push_back(a);
		}
	}

	long long max_weight = 0;
	for (auto& f : floors) {
		long long delta_weight = 0;
		for (auto& in : f.second.inputs) {
			delta_weight = max(delta_weight, floors[in].weight);
		}
		if (f.second.inputs.size())
			delta_weight++;

		f.second.weight += delta_weight;

		max_weight = max(max_weight, f.second.weight);
	}

	cout << max_weight;
}