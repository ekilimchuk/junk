#include <stdlib.h>
#include <stdio.h>
#include <sys/mman.h>
#include <sys/stat.h>
#include <fcntl.h>
#include <unistd.h>
#include <string.h>
#include <openssl/sha.h>
#include <openssl/hmac.h>
#include <twofish.h>

#define TWOFISH_BLOCK_SIZE 16
#define TWOFISH_KEY_SIZE 32
#define KEY_SIZE 32

struct Psafe3 {
	char tag[4];
	char salt[32];
	char iter[4];
	char p[32];
	char b12[32];
	char b34[32];
	char iv[16];
//	char hdr[];
	char db[];
//	char eof;
//	char hmac;
};

void print_hex_debug(char *p, int n)
{
	for (int i = 0; i < n; i++) {
		 printf("%02x ", p[i]);
	}
	printf("\n");
}

int open_psafe3(char *fname, char **addr, int *fsize, char *err)
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
	*fsize = sb.st_size;
	strcpy(err, "");
	return 0;
}

int get_psafe3_data(char **addr, struct Psafe3 **psafe3_data, int *fsize, int *dbsize, char *mac, char *err)
{
	if (*fsize < sizeof(**psafe3_data) + 16 + 32 ) {
		strcpy(err, "It is not PWS3 file - the file size is smaller than the PWS3 header");
		return 1;
	}
	if (strncmp(*addr, "PWS3", 4)) {
		strcpy(err, "It is not PWS3 file");
		return 1;
	}
	if (strncmp(*addr + *fsize - 16 - 32, "PWS3-EOFPWS3-EOF", 16)) {
		strcpy(err, "It is not PWS3 file - nothing matches the end of file with \"PS3-EOFPWS3-EOF\"");
		return 1;
	}
	*dbsize = *fsize - 16 - 32 - sizeof(**psafe3_data);
	*psafe3_data = (void*) *addr;
	memcpy(mac, *addr + *fsize - 32, 32);
	return 0;
}

int stretch_pswd(char *pswd, char *salt, int iter, char *obuf, char *err)
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

	free(pwd_salt);

	for (int i = 0; i < iter; i++) {
		memcpy(tmpbuf, obuf, 32);
		SHA256(tmpbuf, 32, obuf);
	}
	return 0;
}

int check_key(char *key, char *p, char *err)
{
	char obuf[KEY_SIZE];

	SHA256(key, 32, obuf);
	if (strncmp(p, obuf, 32)) {
		strcpy(err, "invalid password");
		return 1;
	}
	return 0;
}

void twofish_ecb(char *key, char *b, int count, char *res)
{
	Twofish_key xkey;
	Twofish_Byte inblock[TWOFISH_BLOCK_SIZE], outblock[TWOFISH_BLOCK_SIZE];

	Twofish_initialise();
	Twofish_prepare_key(key, TWOFISH_KEY_SIZE, &xkey);
	for (int i = 0; i < count; i++) {
		memcpy(inblock, &b[i*TWOFISH_BLOCK_SIZE], TWOFISH_BLOCK_SIZE);
		Twofish_decrypt(&xkey, inblock, outblock);
		memcpy(&res[i * TWOFISH_BLOCK_SIZE], outblock, TWOFISH_BLOCK_SIZE);
	}
}

void str_xor(char *s1, char *s2, int count, char *res)
{
	for (int i = 0; i < count; i++)
		res[i] = s1[i] ^ s2[i];
}

void twofish_cbc(char *iv, char *key, char *b, int count, char *res)
{
	Twofish_key xkey;
	Twofish_Byte inblock[TWOFISH_BLOCK_SIZE], outblock[TWOFISH_BLOCK_SIZE];
	char x[TWOFISH_BLOCK_SIZE];
	memcpy(x, iv, TWOFISH_BLOCK_SIZE);
	Twofish_initialise();
	Twofish_prepare_key(key, TWOFISH_KEY_SIZE, &xkey);
	for (int i = 0; i < count; i++) {
		memcpy(inblock, &b[i * TWOFISH_BLOCK_SIZE], TWOFISH_BLOCK_SIZE);
		Twofish_decrypt(&xkey, inblock, outblock);
		str_xor(outblock, x, TWOFISH_BLOCK_SIZE, &res[i * TWOFISH_BLOCK_SIZE]);
		memcpy(x, inblock, TWOFISH_BLOCK_SIZE);
	}
}

int read_fields(char *data, char key_l[KEY_SIZE], int dsize, char *mac, char *hmac, char *err)
{
	struct item {
		int len;
		char type;
		char data[];
	};
	unsigned int dlen;
	HMAC_CTX ctx;
	struct item *p;
	HMAC_Init(&ctx, key_l, KEY_SIZE, EVP_sha256());
	p = (void*) data;

	while((void*)&p[0] < (void*)&data[0] + dsize) {
		HMAC_Update(&ctx, p->data, p->len);
		if ((5 + p->len) % TWOFISH_BLOCK_SIZE != 0) {
			p = (void*)&p->data[p->len] + TWOFISH_BLOCK_SIZE - (5 + p->len) % TWOFISH_BLOCK_SIZE;
		} else {
			p = (void*)&p->data[p->len];
		}
	}
	HMAC_Final(&ctx, hmac, &dlen);
	HMAC_cleanup(&ctx);
	if (memcmp(mac, hmac, KEY_SIZE)) {
		strcpy(err, "invalid hmac");
		return 1;
	}
	return 0;
}

void print_error(char *err)
{
	fprintf(stderr, "Error: %s!\n", err);
}

int main(int argc, char **argv)
{
	char *addr;
	struct Psafe3 *psafe3_data;
	char err[100];
	char key[KEY_SIZE];
	char key_k[KEY_SIZE];
	char key_l[KEY_SIZE];
	char *res;
	int fsize = 0, dbsize = 0;
	char mac[KEY_SIZE];
	char hmac[KEY_SIZE];

	if (argc < 2) {
		fprintf(stderr, "Run: %s <file.psafe3>\n", argv[0]);
		return 1;
	}
	
	if (open_psafe3(argv[1], &addr, &fsize, err)) {
		print_error(err);
		return 1;
	}

	if (get_psafe3_data(&addr, &psafe3_data, &fsize, &dbsize, mac, err)) {
		print_error(err);
		return 1;
	}

	if (stretch_pswd("bogus12345", psafe3_data->salt, *((int *)psafe3_data->iter), key, err)) {
		print_error(err);
		return 1;
	}

	if (check_key(key, psafe3_data->p, err)) {
		print_error(err);
		return 1;
	}

	twofish_ecb(key, psafe3_data->b12, 2, key_k);

	twofish_ecb(key, psafe3_data->b34, 2, key_l);

	res = malloc(dbsize);

	twofish_cbc(psafe3_data->iv, key_k, psafe3_data->db, dbsize/TWOFISH_BLOCK_SIZE, res);
	if (read_fields(res, key_l, dbsize, mac, hmac, err)) {
		print_error(err);
		return 1;
	}

// debug
/**/	printf("fsize:\n");
	printf("%i\n", fsize);
	printf("\n");
	printf("dbsize:\n");
	printf("%i\n", dbsize);
	printf("\n");
	printf("h(p'):\n");
	print_hex_debug(psafe3_data->p, sizeof(psafe3_data->p));
	printf("\n");
	printf("key:\n");
	print_hex_debug(key, sizeof(key));
	printf("\n");
	printf("k:\n");
	print_hex_debug(key_k, sizeof(key_k));
	printf("\n");
	printf("l:\n");
	print_hex_debug(key_l, sizeof(key_l));
	printf("\n");
	printf("mac:\n");
	print_hex_debug(mac, KEY_SIZE);
	printf("\n");
	printf("hmac:\n");
	print_hex_debug(hmac, KEY_SIZE);
	printf("\n");/**/
/*	for (int i=0; i<dbsize; i++)
		printf("%c", res[i]);
	printf("\n");/**/
	return 0;
}
