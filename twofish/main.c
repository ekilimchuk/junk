#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <twofish.h>

#define MIN_KEYLEN 2

int main(int argc, char **argv)
{
	size_t keylen;
	Twofish_Byte key[32];
	Twofish_key xkey;
	Twofish_Byte inblock[16], outblock[16];

	if (argc < 3) {
		fprintf(stderr, "Run: %s <key> <text>\n", argv[0]);
		return 1;
	}

	memset(key, 0, sizeof(key));

	keylen = strlen(argv[1]);
	if (keylen < MIN_KEYLEN) {
		fprintf(stderr, "Key material too short.\n");
		return 1;
	}
	if (keylen > sizeof(key))
		keylen = sizeof(key);

	Twofish_initialise();
	
	memset(inblock, 0, sizeof(inblock));
	Twofish_prepare_key(key, keylen, &xkey);
	strcpy(inblock, argv[2]);
	printf("original: %s\n", inblock);
	Twofish_encrypt(&xkey, inblock, outblock);
	printf("encrypt: %s\n", outblock);
	
	memset(inblock, 0, sizeof(inblock));
	Twofish_prepare_key(key, keylen, &xkey);
	strcpy(inblock, outblock);
	printf("original: %s\n", inblock);

	Twofish_decrypt(&xkey, inblock, outblock);
	printf("decrypt: %s\n", outblock);

	return 0;
}
