all: main.c
	gcc main.c -ltwofish -o main

test:
	echo -n qwerty1 | ./main -e key1 | ./main -d key1
	echo -n qwerty1 | ./main -e key2 | ./main -d key2
	echo -n qwerty1 | ./main -e key1 | ./main -d key2
	echo -n qwerty2 | ./main -e key1 | ./main -d key1
	echo -n qwerty2 | ./main -e key2 | ./main -d key2
	echo -n qwerty2 | ./main -e key1 | ./main -d key2

clean:
	rm -f main
