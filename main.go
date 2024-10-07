package main

import (
	"fmt"

	"github.com/hamid-nazar/distributed-file-storage/p2p"
)

func main() {
	tcpOptions := p2p.TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    p2p.NOPHandshakefunc,
		Decoder:       &p2p.GOBDecoder{},
	}

	transport := p2p.NewTCPTransport(tcpOptions)

	if err := transport.ListenAndAccept(); err != nil {
		fmt.Printf("ListenAndAccept error: %s\n", err)
	}

	select {}
}
