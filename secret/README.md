# Secret-server
## Build
```
$ make
```

## Clean
```
$ make clean
```

## Gen server rsa key:
```
$ ./secret-gen-keys -private ./keys/server_private_key.pem -public ./keys/server_public_key.pem
```
## Gen client rsa key:
```
$ ./secret-gen-keys -private ./keys/client_private_key.pem -public ./keys/client_public_key.pem
```

## Run server:
```
$ ./secret-server 
```

## Use client:
```
$ ./secret-cli -h
```

