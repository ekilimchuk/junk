#include <stdlib.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <openssl/sha.h>
#include <stdint.h>

struct Psafe3 {
	char tag[4];
	char salt[32];
	char iter[4];
	char p[32];
	char h[256];
	char b1[16];
	char b2[16];
	char b3[16];
	char b4[16];
	char iv[16];
//	char hdr[];
	char db[];
//	char eof;
//	char hmac;
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

int get_psafe3_data(char **addr, struct Psafe3 **psafe3_data, char *err)
{
	if (strncmp(*addr, "PWS3", 4)) {
		strcpy(err, "is non PWS3");
		return 1;
	}
	*psafe3_data = (void*) *addr;
	return 0;
}

int stretch_pswd(char *pswd, char *salt, int32_t iter, char *obuf, char *err)
{
	char tmpbuf[32];

	char *pwd_salt = malloc(strlen(pswd) + 32);
	if (pwd_salt == NULL) {
		strcpy(err, "malloc");
		return 1;
	}

	memcpy(pwd_salt, pswd, strlen(pswd));
	memcpy(pwd_salt + strlen(pswd), salt, 32);

	SHA256(pwd_salt, strlen(pwd_salt), obuf);
	for (long int i = 0; i <= iter; i++) {
		for (int j = 0; j < 32; j++) {
			tmpbuf[j] = obuf[j];
		}
		SHA256(tmpbuf, 32, obuf);
	}
	return 0;
}

void print_error(char *err)
{
	fprintf(stderr, "Error: %s!\n", err);
}

void print_hex_debug(char *p, int n)
{
	for (int i = 0; i < n; i++) {
		 printf("%02x ", p[i]);
	}
	printf("\n");
}

int main(int argc, char **argv)
{
	char *addr;
	struct Psafe3 *psafe3_data;
	size_t size, s;
	char err[100] = "";
	int32_t iter = 0;
	char p[32] = "";

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

	// little-endian
	memcpy(&iter, (*psafe3_data).iter, sizeof((*psafe3_data).iter));

	if (stretch_pswd("bogus12345", (*psafe3_data).salt, iter, p, err)) {
		print_error(err);
		return 1;
	}

	// debug
	printf("PWS3:\n");
	print_hex_debug((*psafe3_data).tag, sizeof((*psafe3_data).tag));
	printf("P:\n");
	print_hex_debug((*psafe3_data).p, sizeof((*psafe3_data).p));
	printf("sha256:\n");
	print_hex_debug(p, sizeof(p));

	return 0;
}
