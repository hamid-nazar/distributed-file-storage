package main

import (
	"fmt"

	"github.com/hamid-nazar/distributed-file-storage/p2p"
)

func main() {
	transport := p2p.NewTCPTransport(":3000")

	if err := transport.ListenAndAccept(); err != nil {
		fmt.Printf("ListenAndAccept error: %s\n", err)
	}
}
