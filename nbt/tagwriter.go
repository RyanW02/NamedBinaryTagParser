package nbt

func (t TagEnd) write(writer *Writer) error {
	return writer.writeTagId(t.Type())
}

func (t TagByte) write(writer *Writer) error {
	return writer.writeUint8(byte(t))
}

func (t TagShort) write(writer *Writer) error {
	return writer.writeInt16(int16(t))
}

func (t TagInt) write(writer *Writer) error {
	return writer.writeInt32(int32(t))
}

func (t TagLong) write(writer *Writer) error {
	return writer.writeInt64(int64(t))
}

func (t TagFloat) write(writer *Writer) error {
	return writer.writeFloat32(float32(t))
}

func (t TagDouble) write(writer *Writer) error {
	return writer.writeFloat64(float64(t))
}

func (t TagByteArray) write(writer *Writer) error {
	if err := writer.writeInt32(int32(len(t))); err != nil {
		return err
	}

	return writer.writeBytes(t)
}

func (t TagString) write(writer *Writer) error {
	return writer.writeUtf8(string(t))
}

func (t TagList) write(writer *Writer) error {
	// write type ID
	if err := writer.writeTagId(t.ElementType); err != nil {
		return err
	}

	// write len
	if err := writer.writeInt32(int32(len(t.Elements))); err != nil {
		return err
	}

	for _, el := range t.Elements {
		if err := el.write(writer); err != nil {
			return err
		}
	}

	return nil
}

func (t TagCompound) write(writer *Writer) error {
	for name, value := range t {
		if err := writer.writeMeta(value, name); err != nil {
			return err
		}

		if err := value.write(writer); err != nil {
			return err
		}
	}

	return TagEnd{}.write(writer)
}

func (t TagIntArray) write(writer *Writer) error {
	// write len
	if err := writer.writeInt32(int32(len(t))); err != nil {
		return err
	}

	for _, value := range t {
		if err := writer.writeInt32(value); err != nil {
			return err
		}
	}

	return nil
}

func (t TagLongArray) write(writer *Writer) error {
	// write len
	if err := writer.writeInt32(int32(len(t))); err != nil {
		return err
	}

	for _, value := range t {
		if err := writer.writeInt64(value); err != nil {
			return err
		}
	}

	return nil
}
