.PHONY: all build_main build_read

all: build_main build_read

build_main: main.c
	@gcc main.c -std=c99 -ltwofish -o main

build_read: read.c
	@gcc read.c -std=c99 -lcrypto -ltwofish -o read

clean:
	@rm -f ./main selftest_[ted].tmp ./read
