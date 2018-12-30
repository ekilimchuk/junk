SERVER_OUT := "secret-server"
CLI_OUT := "secret-cli"
API_OUT := "api/api.pb.go"

.PHONY: all api build_server build_cli

all: build_server build_cli

api/api.pb.go: api/api.proto
	@protoc -I api/ --go_out=plugins=grpc:api api/api.proto

api: api/api.pb.go

build_server: api
	@go build -i -v -o $(SERVER_OUT) ./server

build_cli: api
	@go build -i -v -o $(CLI_OUT) ./cli

clean:
	@rm $(SERVER_OUT) $(CLI_OUT) $(API_OUT)

