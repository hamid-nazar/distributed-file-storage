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

type TCPTransportOptions struct {
	ListenAddress string
	ShakeHands    HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener
	mu       sync.RWMutex
	peers    map[net.Addr]Peer
}

func NewTCPTransport(options TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: options,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	listener, err := net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	t.listener = listener

	go t.startAcceptLoop()

	log.Printf("TCP transport listening on port: %s\n", t.ListenAddress)

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept error: %s\n", err)

			continue
		}

		fmt.Printf("New incoming connection: %+v\n", conn)
		go t.handleConn(conn)
	}
}

type temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewPeer(conn, true)

	if err := t.ShakeHands(peer); err != nil {
		conn.Close()
		fmt.Printf("Handshake error: %s\n", err)
		return
	}

	// Read loop
	msg := &temp{}

	for {

		if err := t.Decoder.Decode(conn, msg); err != nil {
			fmt.Printf("Decode error: %s\n", err)
			continue
		}
	}

}
