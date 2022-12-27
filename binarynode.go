package main

import "encoding/binary"

type BinaryNode struct {
	buffer []byte
	index  int
}

func (b *BinaryNode) ReadU8() uint8 {
	ret := b.buffer[b.index]
	b.index++
	return ret
}

func (b *BinaryNode) ReadU16() uint16 {
	ret := binary.LittleEndian.Uint16(b.buffer[b.index:])
	b.index += 2
	return ret
}

func (b *BinaryNode) ReadU32() uint32 {
	ret := binary.LittleEndian.Uint32(b.buffer[b.index:])
	b.index += 4
	return ret
}

func (b *BinaryNode) ReadU64() uint64 {
	ret := binary.LittleEndian.Uint64(b.buffer[b.index:])
	b.index += 8
	return ret
}

func (b *BinaryNode) Read32() int32 {
	ret := int32(b.buffer[b.index]) | int32(b.buffer[b.index+1])<<8 | int32(b.buffer[b.index+2])<<16 | int32(b.buffer[b.index+3])<<24
	b.index += 4
	return ret
}

func (b *BinaryNode) ReadBool() bool {
	return b.ReadU8() != 0
}

func (b *BinaryNode) Read(len uint16) []byte {
	b.index += int(len)
	return b.buffer[b.index-int(len) : b.index]
}

func (b *BinaryNode) Skip(v uint16) {
	b.index += int(v)
}

func (b *BinaryNode) Empty() bool {
	return int(b.index) >= len(b.buffer)
}
