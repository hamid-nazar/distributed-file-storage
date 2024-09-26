package p2p

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// Peer is the interface that represents the rempte node
// over a TCP established connection
type TCPPeer struct {
	// conn is is the underlying connection of the peer
	conn net.Conn
	// if a connection is retrieved => outBound = true
	// if a connection is accepted => outBound = false
	outBound bool
}

func NewPeer(conn net.Conn, outBound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outBound: outBound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err != nil {
		return err
	}

	t.startAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", t.listenAddress)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewPeer(conn, true)

	fmt.Printf("New incoming connection: %+v\n", peer)
}
