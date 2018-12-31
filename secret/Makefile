SERVER_OUT := "secret-server"
CLI_OUT := "secret-cli"
GEN_KEYS_OUT := "secret-gen-keys"
API_OUT := "api/api.pb.go"

.PHONY: all api build_server build_cli build_gen_keys

all: build_server build_cli build_gen_keys

api/api.pb.go: api/api.proto
	@protoc -I api/ --go_out=plugins=grpc:api api/api.proto

api: api/api.pb.go

build_server: api
	@go build -i -v -o $(SERVER_OUT) ./server

build_cli: api
	@go build -i -v -o $(CLI_OUT) ./cli

build_gen_keys:
	@go build -i -v -o $(GEN_KEYS_OUT) ./gen-keys

clean:
	@rm $(SERVER_OUT) $(CLI_OUT) $(API_OUT) $(GEN_KEYS_OUT)

