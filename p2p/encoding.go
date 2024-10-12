package p2p

import (
	"encoding/gob"
	"io"
)

type Decoder interface {
	Decode(io.Reader, *RPC) error
}

type GOBDecoder struct {
}

func (dec *GOBDecoder) Decode(r io.Reader, rpc *RPC) error {
	return gob.NewDecoder(r).Decode(rpc)
}

type DefaultDecoder struct {
}

func (dec *DefaultDecoder) Decode(r io.Reader, rpc *RPC) error {
	buffer := make([]byte, 1028)
	n, err := r.Read(buffer)

	if err != nil {
		return err
	}

	rpc.Payload = buffer[:n]

	return nil
}
