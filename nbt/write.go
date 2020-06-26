package nbt

import (
	"encoding/binary"
	"math"
)

func (w *Writer) writeBytes(bytes []byte) (err error) {
	_, err = w.writer.Write(bytes)
	return
}

func (w *Writer) writeUint8(i uint8) error {
	return w.writeBytes([]byte{i})
}

func (w *Writer) writeTagId(id TagId) error {
	return w.writeUint8(uint8(id))
}

func (w *Writer) writeUint16(i uint16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, i)
	return w.writeBytes(b)
}

func (w *Writer) writeUint32(i uint32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, i)
	return w.writeBytes(b)
}

func (w *Writer) writeUint64(i uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, i)
	return w.writeBytes(b)
}

func (w *Writer) writeFloat32(f float32) error {
	return w.writeUint32(math.Float32bits(f))
}

func (w *Writer) writeFloat64(f float64) error {
	return w.writeUint64(math.Float64bits(f))
}

func (w *Writer) writeInt8(i int8) error {
	return w.writeUint8(uint8(i))
}

func (w *Writer) writeInt16(i int16) error {
	return w.writeUint16(uint16(i))
}

func (w *Writer) writeInt32(i int32) error {
	return w.writeUint32(uint32(i))
}

func (w *Writer) writeInt64(i int64) error {
	return w.writeUint64(uint64(i))
}

func (w *Writer) writeUtf8(s string) error {
	// write length
	if err := w.writeUint16(uint16(len(s))); err != nil {
		return err
	}

	return w.writeBytes([]byte(s))
}

func (w *Writer) writeMeta(tag Tag, name string) error {
	if err := w.writeTagId(tag.Type()); err != nil {
		return err
	}

	if tag.Type() != TagTypeEnd {
		if err := w.writeUtf8(name); err != nil {
			return err
		}
	}

	return nil
}
