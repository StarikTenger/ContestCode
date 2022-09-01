#include <iostream>
#include <vector>
#include <algorithm>
#include <string>
#include <map>
#include <set>
#include <stack>

using namespace std;

bool is_num(string s) {
	return s[0] > '0' && s[0] < '9' || s[0] == '-';
}

vector<string> split(string str, string delimiter) {
	vector<string> tokens;

	size_t pos = 0;
	std::string token;
	while ((pos = str.find(delimiter)) != std::string::npos) {
		token = str.substr(0, pos);
		tokens.push_back(token);
		str.erase(0, pos + delimiter.length());
	}
	tokens.push_back(str);
	return tokens;
}

struct Context {
	map<string, int> varaibles;
	stack<map<string, int>> saved;

	void save() {
		saved.push({});
	}

	void restore() {
		for (auto& var : saved.top()) {
			varaibles[var.first] = var.second;
		}
		saved.pop();
	}

	void assign(string var1, string var2) {
		if (saved.size()) {
			saved.top()[var1] = varaibles[var1];
		}
		varaibles[var1] = varaibles[var2];
		cout << varaibles[var1] << "\n";
	}

	void assign(string var, int val) {
		if (saved.size()) {
			saved.top()[var] = varaibles[var];
		}
		varaibles[var] = val;
	}
};

int main() {
	Context context;

	std::string line;
	while (std::getline(std::cin, line)) {
		if (line == "{") {
			context.save();
			continue;
		}
		if (line == "}") {
			context.restore();
			continue;
		}

		auto tokens = split(line, "=");
		if (is_num(tokens[1])) {
			context.assign(tokens[0], stoi(tokens[1]));
		}
		else {
			context.assign(tokens[0], tokens[1]);
		}
	}
}