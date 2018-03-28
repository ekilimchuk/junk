#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <twofish.h>

#define MIN_KEYLEN 2
#define BLOCK_SIZE 16
#define KEY_SIZE 32

int main(int argc, char **argv)
{
	size_t keylen;
	Twofish_Byte key[KEY_SIZE];
	Twofish_key xkey;
	Twofish_Byte inblock[BLOCK_SIZE], outblock[BLOCK_SIZE];
	int encrypt = 1;	

	if (argc < 3) {
		fprintf(stderr, "Run: %s <-e|-d> <key>\n", argv[0]);
		return 1;
	}

	if (strcmp(argv[1], "-d") == 0)
		encrypt = 0;

	memset(key, 0, sizeof(key));

	keylen = strlen(argv[2]);
	if (keylen < MIN_KEYLEN) {
		fprintf(stderr, "Key material too short.\n");
		return 1;
	}
	if (keylen > sizeof(key))
		keylen = sizeof(key);

	Twofish_initialise();
	strncpy((char *) key, argv[2], sizeof(key));	
	memset(inblock, 0, sizeof(inblock));
	memset(outblock, 0, sizeof(outblock));
	Twofish_prepare_key(key, keylen, &xkey);

	while (read(STDIN_FILENO, inblock, sizeof(inblock)) > 0) {
		if (encrypt) {
			Twofish_encrypt(&xkey, inblock, outblock);
			write(STDOUT_FILENO, outblock, sizeof(outblock));
		} else {	
			Twofish_decrypt(&xkey, inblock, outblock);
			for (int i = 0; i < BLOCK_SIZE && outblock[i] != '\0'; i++)
				printf("%c", outblock[i]);
		}
		memset(inblock, 0, sizeof(inblock));
		memset(outblock, 0, sizeof(outblock));
	}
	return 0;
}
