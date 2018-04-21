all: target

BIN_PATH=`pwd`/bin/
GO_PATH=`pwd`

network: network/
	GOPATH=$(GO_PATH) go build network
common: common/
	GOPATH=$(GO_PATH) go build common

target: src/server.go 
	GOPATH=$(GO_PATH) GOBIN=$(BIN_PATH) go install src/server.go
	# GOPATH=f:\goworkspace\GoServer\ GOBIN=f:\goworkspace\GoServer\bin\ go install src\server.go
