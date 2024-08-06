package frame

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

type FramePayload []byte
type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error
	Decode(io.Reader) (FramePayload, error)
}

var ErrShortWrite = errors.New("short write")
var ErrShortRead = errors.New("short read")

type myFrameCodec struct{}

func (p *myFrameCodec) Encode(w io.Writer, payload FramePayload) error {
	var f = payload
	var totalLen = int32(len(payload)) + 4
	err := binary.Write(w, binary.BigEndian, &totalLen)
	if err != nil {
		return err
	}
	n, err := w.Write([]byte(f))
	if err != nil {
		return err
	}
	if n != len(f) {
		return ErrShortWrite
	}
	return nil
}

func (p *myFrameCodec) Decode(r io.Reader) (FramePayload, error) {
	fmt.Println("Decode start")
	var totalLen int32
	err := binary.Read(r, binary.BigEndian, &totalLen)
	fmt.Println("Decode end")
	if err != nil {
		return nil, err
	}
	buf := make([]byte, totalLen-4)
	n, err := io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}
	if n != int(totalLen-4) {
		return nil, ErrShortRead
	}

	return FramePayload(buf), nil
}

func NewMyFrameCodec() StreamFrameCodec {
	return &myFrameCodec{}
}
