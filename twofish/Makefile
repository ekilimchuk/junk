all: main.c
	gcc main.c -ltwofish -o main

test:
	echo
	./main key1 qwerty1
	echo
	./main key2 qwerty2
	echo
	./main key3 qwerty3

clean:
	rm -f main
