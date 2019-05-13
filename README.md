# udp-proxy

This is a simple utility tool for proxying udp packets to other listening clients.

### Install
```sh
go get -u github.com/kvartborg/udp-proxy
```

### Usage
```sh
udp-proxy <host:port> <destination:port> <...>

udp-proxy localhost:3000 localhost:4000 localhost:5000
```
