package p2p

import (
	"fmt"
	"log"
	"net"
	"strings"
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

func (p *TCPPeer) Close() error {
	return p.conn.Close()
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
	OnPeer        func(Peer) error
}

type TCPTransport struct {
	TCPTransportOptions
	listener   net.Listener
	rpcChannel chan RPC
}

func NewTCPTransport(options TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: options,
		rpcChannel:          make(chan RPC),
	}
}

// Consume implements the Transport interface, which will return read-only channel
// for reading incoming messages received from another peer in the network
func (t *TCPTransport) Consume() <-chan RPC {
	return t.rpcChannel
}

func (t *TCPTransport) ListenAndAccept() error {
	var err error

	TCPListener, err := net.Listen("tcp", t.ListenAddress)
	if err != nil {
		return err
	}

	t.listener = TCPListener

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

func (t *TCPTransport) handleConn(conn net.Conn) {
	var err error

	defer func() {
		fmt.Printf("Dropping peer connection: %s\n", err)
		conn.Close()
	}()

	peer := NewPeer(conn, true)

	if err = t.ShakeHands(peer); err != nil {
		return
	}

	if t.OnPeer != nil {
		if err = t.OnPeer(peer); err != nil {
			return
		}
	}

	// Read loop

	for {

		rpc := RPC{}

		if err := t.Decoder.Decode(conn, &rpc); err != nil {
			fmt.Printf("Decoding error: %s\n", err)
			return
		}

		message := strings.TrimSpace(string(rpc.Payload))

		if message == "exit" {
			fmt.Println("Received exit command, closing connection...")

			conn.Write([]byte("Bye :)\n"))
			conn.Close()
			return
		}

		rpc.From = conn.RemoteAddr()

		fmt.Printf("Received message: %s\n", rpc)

		// Write back to the client

		// response := "You sent: " + string(rpc.Payload) + "\n"
		// conn.Write([]byte(response))

		t.rpcChannel <- rpc
	}

}
