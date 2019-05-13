# udp-proxy

[![Build Status](https://travis-ci.com/kvartborg/udp-proxy.svg?branch=master)](https://travis-ci.com/kvartborg/udp-proxy)
[![GoDoc](https://godoc.org/github.com/kvartborg/udp-proxy?status.svg)](https://godoc.org/github.com/kvartborg/udp-proxy)
[![Go Report Card](https://goreportcard.com/badge/github.com/kvartborg/udp-proxy)](https://goreportcard.com/report/github.com/kvartborg/udp-proxy)

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
