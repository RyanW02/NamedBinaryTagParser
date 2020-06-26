package nbt

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
)

type Parser struct {
	reader          io.Reader
}

func NewParser(reader io.Reader) *Parser {
	return &Parser{
		reader: reader,
	}
}

var ErrNotCompound = errors.New("first tag must be a compound tag")

func (p *Parser) Read() (tag TagCompound, outerName string, err error) {
	var compressionType CompressionType
	compressionType, err = p.getCompressionType()
	if err != nil {
		return
	}

	if err = p.swapReader(compressionType); err != nil {
		return
	}

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

func (p *Parser) getCompressionType() (CompressionType, error) {
	// convert to bufio reader so we can peak
	bytes, err := bufio.NewReader(p.reader).Peek(1)
	if err != nil {
		return Uncompressed, err
	}

	head := bytes[0]
	switch head {
	case 0x1f:
		return Gzip, nil
	case 0x78:
		return Zlib, nil
	default:
		return Uncompressed, nil
	}
}

func (p *Parser) swapReader(compressionType CompressionType) (err error) {
	var reader io.Reader

	switch compressionType {
	case Gzip:
		reader, err = gzip.NewReader(p.reader)
	case Zlib:
		reader, err = zlib.NewReader(p.reader)
	}

	if reader != nil {
		p.reader = bufio.NewReader(reader)
	}

	return err
}
