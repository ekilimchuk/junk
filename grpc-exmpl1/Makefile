SERVER_OUT := "myserver"
CLIENT_OUT := "myclient"
API_OUT := "api/api.pb.go"

.PHONY: all api build_server build_client

all: build_server build_client

api/api.pb.go: api/api.proto
	@protoc -I api/ --go_out=plugins=grpc:api api/api.proto

api: api/api.pb.go

build_server: api
	@go build -i -v -o $(SERVER_OUT) ./server

build_client: api
	@go build -i -v -o $(CLIENT_OUT) ./client

clean:
	@rm $(SERVER_OUT) $(CLIENT_OUT) $(API_OUT)
