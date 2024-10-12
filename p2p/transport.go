package p2p

// Peer is the interface that represents the rempte node.
type Peer interface {
	Close() error
}

// Transport is anything that handles network communication between nodes.
// It is responsible for dialing and listening for incoming connections.
type Transport interface {
	// Start starts the transport.
	ListenAndAccept() error
	Consume() <-chan RPC
}
