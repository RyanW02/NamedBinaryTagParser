package nbt

import (
	"io"
)

type Writer struct {
	writer io.Writer
}

func NewWriter(writer io.Writer) *Writer {
	return &Writer{
		writer: writer,
	}
}

func (w *Writer) Write(tag TagCompound, outerName string) (err error) {
	// writer compound meta
	if err := w.writeMeta(tag, outerName); err != nil {
		return err
	}

	return tag.write(w)
}
