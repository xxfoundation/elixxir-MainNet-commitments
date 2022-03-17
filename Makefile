.PHONY: update master release update_master update_release build clean

clean:
	rm -rf vendor/
	go mod vendor

update:
	-GOFLAGS="" go get -u all

build:
	go build ./...
	go mod tidy

update_release:
	GOFLAGS="" go get gitlab.com/xx_network/primitives@release
	GOFLAGS="" go get gitlab.com/xx_network/crypto@release

update_master:
	GOFLAGS="" go get gitlab.com/xx_network/primitives@master
	GOFLAGS="" go get gitlab.com/xx_network/crypto@master

master: update_master clean build

release: update_release clean build

linux_server:
	GOOS=linux GOARCH=amd64 go build -o commitments-server.binary server.go

mac_server:
	GOOS=darwin GOARCH=amd64 go build -o commitments-server.binary server.go

client:
	GOOS=js GOARCH=wasm go build -o main.wasm

cli_client_mac:
	GOOS=darwin GOARCH=amd64 go build -o commitments-client.binary client.go

cli_client_linux:
	GOOS=linux GOARCH=amd64 go build -o commitments-client.binary client.go
