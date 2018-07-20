all: gateway game 

CLIENT_PROTO=$(wildcard ./protoc/*.proto)

proto:
	protoc --go_out=./src/proto/ -I=./protoc/ $(CLIENT_PROTO)

gateway: proto
	go install src/gateway

game: proto
	go install src/game

.PHONY: clean

clean:
	rm ./protoc/*.go