package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOptions := TCPTransportOptions{
		ListenAddress: ":4000",
		ShakeHands:    NOPHandshakefunc,
		Decoder:       &DefaultDecoder{},
	}

	transport := NewTCPTransport(tcpOptions)

	assert.Equal(t, transport.ListenAddress, ":4000")

	assert.Nil(t, transport.listener)

}
