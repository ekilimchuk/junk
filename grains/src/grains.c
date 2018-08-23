#define MIN 1
#define MAX 64
unsigned long square(const int n) {
	if ( n < MIN || n > MAX) {
		return 0ul;
	}
	return 1ul << (n - 1);
}

unsigned long total() {
	unsigned long sum = 0ul;
	for (int i = MIN; i <= MAX; i++) {
		sum += square(i);
	}
	return sum;
}
