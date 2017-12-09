// https://programmers.co.kr/learn/challenge_codes/147

#include <iostream>

using namespace std;

long long fibonacci(int n) {
	int a, b, c, i;
	a = 0;
	b = 1;
	c = a + b;
	for (i = 0; ; i++) {
		a = b;
		b = c;
		if (b > n) { break; }
		c = a + b;
	}
	return i;
}

int main() {
	int testCase = 47; // 8
	cout <<  fibonacci(testCase) <<  endl;

	return 0;
}
