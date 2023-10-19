package frame

import (
	"encoding/binary"
	"errors"
	"io"
)

type FramePayload []byte

type StreamFrameCodec interface {
	Encode(io.Writer, FramePayload) error
	Decode(io.Reader) (FramePayload, error)
}

var ErrShortRead = errors.New("short read")

var ErrShortWrite = errors.New("short write")

type MyFrameCodec struct {
}

func (m *MyFrameCodec) Encode(w io.Writer, f FramePayload) error {
	l := len(f)
	var total int32 = int32(l) + 4
	err := binary.Write(w, binary.BigEndian, &total)
	if err != nil {
		return err
	}

	n, err := w.Write([]byte(f))
	if err != nil {
		return err
	}
	if n != l {
		return ErrShortWrite
	}

	return nil
}

func (m *MyFrameCodec) Decode(reader io.Reader) (FramePayload, error) {

	var total int32
	err := binary.Read(reader, binary.BigEndian, &total)
	if err != nil {
		return nil, err
	}
	payload := make(FramePayload, total-4)
	n, err := io.ReadFull(reader, payload) //io.ReadFull一般会读满你所需的字节数，除非遇到EOF或ErrUnexpectedEOF。
	if err != nil {
		return nil, err
	}
	if n != int(total)-4 {
		return nil, ErrShortRead
	}

	return payload, nil
}
