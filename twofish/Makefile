all: main.c
	gcc main.c -std=c99 -ltwofish -o main

test:
	openssl rand -base64 500 > selftest_t.tmp
	./main -e key1 > selftest_e.tmp < selftest_t.tmp
	./main -d key1 > selftest_d.tmp < selftest_e.tmp
	diff selftest_t.tmp selftest_d.tmp
	md5sum selftest_t.tmp selftest_d.tmp

clean:
	rm -f ./main selftest_[ted].tmp
