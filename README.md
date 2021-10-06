# MainNet commitments

### Build instructions
The makefile contains three targets for default builds
To build server binaries, use either of the following:

`make linux_server`
`make mac_server`

To compile the webassembly bindings use:

`make client`

### Example config
```yaml
keyPath: ""
certPath: ""
port: ""

# Database connection information
dbUsername: "cmix"
dbPassword: ""
dbName: "cmix_server"
dbAddress: ""
```
