.PHONY: linux_server mac_server client

linux_server:
	GOOS=linux GOARCH=amd64 go build -o commitments-server.binary

mac_server:
	GOOS=darwin GOARCH=amd64 go build -o commitments-server.binary

client:
	GOOS=js GOARCH=wasm go build -o main.wasm
