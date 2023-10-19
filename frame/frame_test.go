package frame

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

func TestMyFrameCodec_Decode(t *testing.T) {

	myCodec := &MyFrameCodec{}

	b := []byte{0x0, 0x0, 0x0, 0x8, 'k', 'k', 'l', 'v'}
	bb := bytes.NewReader(b)
	payload, err := myCodec.Decode(bb)
	if err != nil {
		t.Errorf("Decode happened error: %+v, %+v", err, payload)
	}

	if string(payload) != "kklv" {
		t.Error("decode failed, total -> get:", string(payload), " want:", "kklv")
	}
}

func TestMyFrameCodec_Encode(t *testing.T) {

	b := make([]byte, 0, 20)
	bb := bytes.NewBuffer(b)

	myCodec := &MyFrameCodec{}
	err := myCodec.Encode(bb, []byte("kklv"))
	if err != nil {
		t.Errorf("encode happened error: %+v", err)
	}

	var total int32
	err = binary.Read(bb, binary.BigEndian, &total)
	if err != nil {
		t.Errorf("read happened error: %+v", err)
	}

	if total != 8 {
		t.Error("encode failed, total -> get:", total, " want:", 4)
	}

	payload := make([]byte, total-4)
	_, err = io.ReadFull(bb, payload)
	if err != nil {
		t.Errorf("ReadFull happened error: %+v", err)
	}

	if string(payload) != "kklv" {
		t.Error("encode failed, total -> get:", string(payload), " want:", "kklv")
	}

}
