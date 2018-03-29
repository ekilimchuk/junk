#include <stdlib.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>

int main(int argc, char **argv)
{
	int fd;
	ssize_t s;
	struct stat sb;
	char *addr;
	off_t pa_offset;

	if (argc < 2) {
		fprintf(stderr, "Run: %s <file.psafe3>\n", argv[0]);
		return 1;
	}

	fd = open(argv[1], O_RDONLY);
	if (fd == -1) {
		fprintf(stderr, "Error: open!\n", argv[0]);
		return 1;
	}
	
	if (fstat(fd, &sb) == -1) {
		fprintf(stderr, "Error: fstat!\n", argv[0]);
		return 1;
	}

	addr = mmap(NULL, sb.st_size, PROT_READ, MAP_PRIVATE, fd, 0);
	if (addr == MAP_FAILED) {
		fprintf(stderr, "Error: mmap!\n", argv[0]);
		return 1;
	}

	s = write(STDOUT_FILENO, addr, sb.st_size);
	if (s != sb.st_size) {
		if (s == -1) {
			fprintf(stderr, "Error: write!\n", argv[0]);
			return 1;
		}
	}
	return 0;
}
