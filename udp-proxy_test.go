package main

import (
	"net"
	"testing"
	"time"
)

func TestUDPProxy(t *testing.T) {
	receiver, _ := net.ListenPacket("udp", "127.0.0.1:50001")
	defer receiver.Close()
	go start("127.0.0.1:50000", "127.0.0.1:50001")
	sender, _ := net.Dial("udp", "127.0.0.1:50000")
	defer sender.Close()

	result := make(chan []byte)

	go func() {
		b := make([]byte, 5)
		receiver.ReadFrom(b)
		result <- b
	}()

	sender.Write([]byte("Hello"))

	select {
	case data := <-result:
		if string(data) != "Hello" {
			t.Errorf(
				"received packet did not contain expected message %s; got %s",
				"Hello", string(data),
			)
		}
	case <-time.After(time.Second):
		t.Error("Did not receive udp packet within given timeout")
	}
}
