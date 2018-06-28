server: src/server.go 
	go install src/server.go

CLIENT_PROTO=$(wildcard ./protoc/*.proto)

proto:
	protoc --go_out=./src/proto/ -I=./protoc/ $(CLIENT_PROTO)

.PHONY: clean proto

clean:

