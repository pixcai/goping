package main

import (
	"bytes"
	"encoding/binary"
)

type Packet struct {
	Type        uint8
	Code        uint8
	Checksum    uint16
	Identifier  uint16
	SequenceNum uint16
}

func (p *Packet) Bytes() []byte {
	var buf bytes.Buffer

	binary.Write(&buf, binary.BigEndian, p)
	p.Checksum = Checksum(buf.Bytes())
	buf.Reset()
	binary.Write(&buf, binary.BigEndian, p)

	return buf.Bytes()
}

func NewPacket() *Packet {
	return &Packet{Type: 8, Code: 0, Checksum: 0, Identifier: 0, SequenceNum: 0}
}

func Checksum(data []byte) uint16 {
	var (
		i   int    = 0
		j   int    = len(data)
		sum uint32 = 0
	)

	for j > 1 {
		sum += uint32(data[i])<<8 + uint32(data[i+1])
		i += 2
		j -= 2
	}
	if j > 0 {
		sum += uint32(data[i])
	}
	sum += (sum >> 16)

	return uint16(^sum)
}
