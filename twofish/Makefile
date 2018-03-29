all: main read

main: main.c
	gcc main.c -std=c99 -ltwofish -o main

read: read.c
	gcc read.c -std=c99 -o read

test:
	openssl rand -base64 500 > selftest_t.tmp
	./main -e key1 > selftest_e.tmp < selftest_t.tmp
	./main -d key1 > selftest_d.tmp < selftest_e.tmp
	diff selftest_t.tmp selftest_d.tmp
	md5sum selftest_t.tmp selftest_d.tmp
	./read read.c | wc -l

clean:
	rm -f ./main selftest_[ted].tmp ./read
