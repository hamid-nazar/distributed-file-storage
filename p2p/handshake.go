package p2p

import "errors"

// ErrInvalidHandshake is returned when the handshake between the remote and
// the local node cannot be established
var ErrInvalidHandshake = errors.New("invalid handshake")

// HandshakeFunc is used to verify the remote peer's handshake
type HandshakeFunc func(Peer) error

// NOPHandshakefunc is a HandshakeFunc that does nothing
func NOPHandshakefunc(Peer) error {
	return nil
}
