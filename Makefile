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
