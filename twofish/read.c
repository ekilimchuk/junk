#include <stdlib.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>

struct Psafe3 {
	char *tag;
	char *salt;
	char *p;
	char *iter;
	char *h;
	char *b1;
	char *b2;
	char *b3;
	char *b4;
	char *iv;
	char *hdr;
	char *db;
	char *eof;
	char *hmac;
};

int open_psafe3(char *fname, char **addr, char *err)
{
	int fd;
	struct stat sb;
	fd = open(fname, O_RDONLY);
	if (fd == -1) {
		strcpy(err, "open");
		printf("%s\n", err);
		return 1;
	}
	if (fstat(fd, &sb) == -1) {
		strcpy(err, "fstat");
		return 1;
	}
	*addr = mmap(NULL, sb.st_size, PROT_READ, MAP_PRIVATE, fd, 0);
	if (*addr == MAP_FAILED) {
		strcpy(err, "mmap");
		return 1;
	}
	strcpy(err, "");
	return 0;
}

int get_psafe3_data(char **addr, struct Psafe3 *psafe3_data, char *err)
{
	if (strncmp(*addr, "PWS3", 4)) {
		strcpy(err, "is non PWS3");
		return 1;
	}
	(*psafe3_data).tag = *addr;
	(*psafe3_data).salt = (*psafe3_data).tag + 4;
	(*psafe3_data).p = (*psafe3_data).salt + 32;
	return 0;
}

void print_error(char *err)
{
	fprintf(stderr, "Error: %s!\n", err);
}

int main(int argc, char **argv)
{
	char *addr;
	struct Psafe3 psafe3_data;
	size_t size, s;
	char err[100] = "";

	if (argc < 2) {
		fprintf(stderr, "Run: %s <file.psafe3>\n", argv[0]);
		return 1;
	}
	
	if (open_psafe3(argv[1], &addr, err)) {
		print_error(err);
		return 1;
	}
	if (get_psafe3_data(&addr, &psafe3_data, err)) {
		print_error(err);
		return 1;
	}

	write(STDOUT_FILENO, psafe3_data.tag, 4);
	write(STDOUT_FILENO, psafe3_data.salt, 32);
	write(STDOUT_FILENO, psafe3_data.p, 256);

	return 0;
}
