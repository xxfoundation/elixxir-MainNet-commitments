# MainNet commitments

### Client usage
```
./client.binary -k ../wrapper/creds/1/cmix-key.key -i ../wrapper/creds/1/cmix-IDF.json -n "NOMINATOR WALLET" -v "VALIDATOR WALLET"
```

### Build instructions
The makefile contains three targets for default builds
To build server binaries, use either of the following:

`make linux_server`
`make mac_server`

To build the CLI binary use:
`make cli_client_mac`
`make cli_client_linux`

To compile the webassembly bindings use:

`make client`

### Example config
```yaml
# ==================================
# Commitments Server Configuration
# ==================================

# START YAML ===
# Verbose logging
logLevel: 1
# Path to log file
log: "/cmix/commitments.log"

# Database connection information
dbUsername: ""
dbPassword: ""
dbName: ""
dbAddress: "0.0.0.0:5432"

# Path to this server's private key file
keyPath: "~/.elixxir/commitments.elixxir.io.key"
# Path to this server's certificate file
certPath: "~/.elixxir/commitments.elixxir.io.crt"
# The listening port of this server
port: 11420

# Hash of valid contract
contractHash: "bcMp4W4W3Gd1rO56QUt54Cfr_AEBQbkIcxMIRH1PJqh8mXnCYM4bIdjHdMRjD2r-lrewwPVBIsHYeXT6Knopwg=="

# === END YAML

```
