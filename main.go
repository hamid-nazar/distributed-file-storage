package main

import (
	"fmt"

	"github.com/hamid-nazar/distributed-file-storage/p2p"
)

func OnPeer(peer p2p.Peer) error {
	// return fmt.Errorf("OnPeer function not implemented")
	peer.Close()
	return nil
}
func main() {
	tcpOptions := p2p.TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    p2p.NOPHandshakefunc,
		Decoder:       &p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}

	transport := p2p.NewTCPTransport(tcpOptions)

	go func() {

		for {
			rpc := <-transport.Consume()
			fmt.Printf("Received message from peer: %s\n", rpc)
		}
	}()

	if err := transport.ListenAndAccept(); err != nil {
		fmt.Printf("ListenAndAccept error: %s\n", err)
	}

	select {}
}
