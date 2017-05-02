package bytes

import (
	"encoding/binary"
)

// BinaryBuffer is a buffer that marshals binary data from a byte slice to
// other Go types and back.
//
// This could be achieved with a bytes.Buffer, but not in a zero-copy way
// because bytes.Buffer and binary.ByteOrder do not interact nicely.
type BinaryBuffer struct {
	buf      []byte
	overflow bool
	order    binary.ByteOrder
}

func NewBinaryBuffer(data []byte, order binary.ByteOrder) *BinaryBuffer {
	return &BinaryBuffer{data, false, order}
}

func (bb *BinaryBuffer) advance(n int) ([]byte, bool) {
	if n > len(bb.buf) {
		bb.overflow = true
		return nil, false
	}
	p := bb.buf[:n]
	bb.buf = bb.buf[n:]
	return p, true
}

func (bb *BinaryBuffer) grow(n int) []byte {
	bb.buf = append(bb.buf, make([]byte, n)...)
	return bb.buf[len(bb.buf)-n:]
}

func (bb *BinaryBuffer) Overflowed() bool {
	return bb.overflow
}

func (bb *BinaryBuffer) Read8() uint8 {
	sl, ok := bb.advance(1)
	if !ok {
		return 0
	}
	return uint8(sl[0])
}

func (bb *BinaryBuffer) Read16() uint16 {
	sl, ok := bb.advance(2)
	if !ok {
		return 0
	}
	return bb.order.Uint16(sl)
}

func (bb *BinaryBuffer) Read32() uint32 {
	sl, ok := bb.advance(4)
	if !ok {
		return 0
	}
	return bb.order.Uint32(sl)
}

func (bb *BinaryBuffer) Read64() uint64 {
	sl, ok := bb.advance(8)
	if !ok {
		return 0
	}
	return bb.order.Uint64(sl)
}

func (bb *BinaryBuffer) ReadSlice() []byte {
	l := bb.Read32()
	sl, _ := bb.advance(int(l))
	return sl
}

func (bb *BinaryBuffer) ReadString() string {
	return string(bb.ReadSlice())
}

func (bb *BinaryBuffer) Write8(v uint8) {
	bb.grow(1)[0] = byte(v)
}

func (bb *BinaryBuffer) Write16(v uint16) {
	bb.order.PutUint16(bb.grow(2), v)
}

func (bb *BinaryBuffer) Write32(v uint32) {
	bb.order.PutUint32(bb.grow(4), v)
}

func (bb *BinaryBuffer) Write64(v uint64) {
	bb.order.PutUint64(bb.grow(8), v)
}
