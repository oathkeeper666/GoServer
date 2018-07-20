all: gateway game

CLIENT_PROTO=$(wildcard ./protoc/*.proto)

proto:
	protoc --go_out=./src/proto/ -I=./protoc/ $(CLIENT_PROTO)

gateway: proto
	go install gateway

game: proto
	go install game

.PHONY: clean gateway game

clean:
	rm ./src/proto/*.go
