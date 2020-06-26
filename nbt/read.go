package nbt

import (
	"encoding/binary"
	"io"
	"math"
)

func (p *Parser) readBytes(n int) ([]byte, error) {
	bytes := make([]byte, n)

	read, err := p.reader.Read(bytes)
	if err != nil {
		return nil, err
	}

	if read != n {
		return nil, io.ErrUnexpectedEOF
	}

	return bytes, nil
}

func (p *Parser) readUint8() (uint8, error) {
	bytes, err := p.readBytes(8 >> 3)
	if err != nil {
		return 0, nil
	}

	return bytes[0], nil
}

func (p *Parser) readTagId() (TagId, error) {
	id, err := p.readUint8()
	if err != nil {
		return 0, err
	}

	return TagId(id), err
}

func (p *Parser) readUint16() (uint16, error) {
	bytes, err := p.readBytes(16 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint16(bytes), nil
}

func (p *Parser) readUint32() (uint32, error) {
	bytes, err := p.readBytes(32 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint32(bytes), nil
}

func (p *Parser) readUint64() (uint64, error) {
	bytes, err := p.readBytes(64 >> 3)
	if err != nil {
		return 0, nil
	}

	return binary.BigEndian.Uint64(bytes), nil
}

func (p *Parser) readFloat32() (float32, error) {
	bytes, err := p.readUint32()
	if err != nil {
		return 0, nil
	}

	return math.Float32frombits(bytes), nil
}

func (p *Parser) readFloat64() (float64, error) {
	bytes, err := p.readUint64()
	if err != nil {
		return 0, nil
	}

	return math.Float64frombits(bytes), nil
}

func (p *Parser) readInt8() (int8, error) {
	unsigned, err := p.readUint8()
	if err != nil {
		return 0, err
	}

	return int8(unsigned), err
}

func (p *Parser) readInt16() (int16, error) {
	unsigned, err := p.readUint16()
	if err != nil {
		return 0, err
	}

	return int16(unsigned), err
}

func (p *Parser) readInt32() (int32, error) {
	unsigned, err := p.readUint32()
	if err != nil {
		return 0, err
	}

	return int32(unsigned), err
}

func (p *Parser) readInt64() (int64, error) {
	unsigned, err := p.readUint64()
	if err != nil {
		return 0, err
	}

	return int64(unsigned), err
}

func (p *Parser) readUtf8(length int) (string, error) {
	bytes, err := p.readBytes(length)
	if err != nil {
		return "", nil
	}

	return string(bytes), nil
}

