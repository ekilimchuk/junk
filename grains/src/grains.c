#define MIN 1
#define MAX 64
unsigned long long square(const unsigned int n) {
	if ( n < MIN || n > MAX) {
		return 0ull;
	}
	return 1ull << (n - 1);
}

unsigned long long total() {
	unsigned long long sum = 0ull;
	for (int i = MIN; i <= MAX; i++) {
		sum += square(i);
	}
	return sum;
}
