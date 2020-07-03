package nbt

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
)

type Parser struct {
	reader io.Reader
}

func NewParser(reader io.Reader) (*Parser, error) {
	var buf bytes.Buffer
	_, err := buf.ReadFrom(reader)
	if err != nil {
		return nil, err
	}

	compressionType, err := getCompressionType(buf.Bytes()[0])
	if err != nil {
		return nil, err
	}

	compressionReader, err := getReader(&buf, compressionType)
	if err != nil {
		return nil, err
	}

	return &Parser{
		reader: compressionReader,
	}, nil
}

var ErrNotCompound = errors.New("first tag must be a compound tag")

func (p *Parser) Read() (tag TagCompound, outerName string, err error) {
	// read first tag id
	var tagId TagId
	tagId, err = p.readTagId()
	if err != nil {
		return
	}

	if tagId != TagTypeCompound {
		err = ErrNotCompound
		return
	}

	// parse tag name
	var nameLength uint16
	nameLength, err = p.readUint16()
	if err != nil {
		return
	}

	outerName, err = p.readUtf8(int(nameLength))
	if err != nil {
		return
	}

	var outer Tag
	outer, err = tagId.Read(p)

	tag = outer.(TagCompound)

	return
}

func getCompressionType(head byte) (CompressionType, error) {
	switch head {
	case 0x1f:
		return Gzip, nil
	case 0x78:
		return Zlib, nil
	default:
		return Uncompressed, nil
	}
}

func getReader(existing io.Reader, compressionType CompressionType) (reader io.Reader, err error) {
	switch compressionType {
	case Gzip:
		reader, err = gzip.NewReader(existing)
	case Zlib:
		reader, err = zlib.NewReader(existing)
	default:
		return existing, nil
	}

	return
}
