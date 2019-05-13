package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	listener    net.PacketConn
	connections []net.Conn
	packetSize  *int
	help        *bool
)

func init() {
	packetSize = flag.Int("packet-size", 512, "Size of packets in bytes")
	flag.Parse()
}

func start(host string, destinations ...string) {
	var err error
	listener, err = net.ListenPacket("udp", host)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	for _, destination := range destinations {
		if destination[0] == '-' {
			continue
		}

		nc, err := net.Dial("udp", destination)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		connections = append(connections, nc)
	}

	b := make([]byte, *packetSize)
	for {
		n, _, err := listener.ReadFrom(b)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}

		for _, connection := range connections {
			_, err := connection.Write(b[:n])

			if err != nil {
				fmt.Fprint(os.Stderr, err)
				os.Exit(1)
			}
		}
	}
}

func close() {
	listener.Close()
	for _, conn := range connections {
		conn.Close()
	}
}

func main() {
	var (
		host         string
		destinations []string
	)

	if len(os.Args) < 3 {
		fmt.Fprint(
			os.Stderr,
			"udp-proxy missing <host> and at least one <distination>\n\nExample:\n\tudp-proxy localhost:3000 localhost:4000",
		)
		os.Exit(1)
	}

	if os.Args[1][0] == '-' {
		host = os.Args[2]
		destinations = os.Args[3:]
	} else {
		host = os.Args[1]
		destinations = os.Args[2:]
	}

	start(host, destinations...)
	defer func() {
		close()
	}()
}
