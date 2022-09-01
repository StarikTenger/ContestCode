#include <iostream>
#include <vector>
#include <string>
#include <map>
#include <algorithm>
#include <numeric>

using namespace std;

long long prime_factors(long long n) {
	long long maxf = 1;

	while (n % 2 == 0) {
		maxf = 2;
		n = n / 2;
	}

	for (long long i = 3; i * i <= n; i = i + 2) {
		while (n % i == 0) {
			maxf = i;
			n = n / i;
		}
	}

	if (n > 2)
		maxf = n;

	return maxf;
}

long long gcd(long long a, long long b) {
	while (b) {
		a %= b;
		swap(a, b);
	}
	return a;
}

int main() {
	long long t;
	cin >> t;
	
	for (long long i = 0; i < t; i++) {
		long long a, b;
		cin >> a >> b;
		auto x = gcd(a, b);
		a /= x; b /= x;

		auto max_a = x * prime_factors(a);
		auto max_b = x * prime_factors(b);

		if (max_a > max_b)
			cout << max_a;
		else
			cout << max_b;
		cout << "\n";
	}
	
}