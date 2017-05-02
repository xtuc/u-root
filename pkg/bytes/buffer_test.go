package bytes

import (
	"encoding/binary"
	"testing"
)

type foo struct {
	a uint8
	b uint16
	c uint32
	d uint64
}

func (f foo) Marshal(b *BinaryBuffer) {
	b.Write8(f.a)
	b.Write16(f.b)
	b.Write32(f.c)
	b.Write64(f.d)
}

func (f *foo) Unmarshal(b *BinaryBuffer) {
	f.a = b.Read8()
	f.b = b.Read16()
	f.c = b.Read32()
	f.d = b.Read64()
}

func TestBinaryBuffer(t *testing.T) {
	buf := NewBinaryBuffer([]byte{}, binary.LittleEndian)

	want := foo{1, 1 << 5, 1 << 9, 1 << 13}
	want.Marshal(buf)

	got := &foo{}
	got.Unmarshal(buf)
	if want != *got {
		t.Errorf("Encoding/Decoding failed: got %v, want %v", got, want)
	}
}
