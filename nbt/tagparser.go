package nbt

func (p *Parser) readTagByte() (TagByte, error) {
	byte, err := p.readInt8()
	return TagByte(byte), err
}

func (p *Parser) readTagShort() (TagShort, error) {
	i, err := p.readInt16()
	return TagShort(i), err
}

func (p *Parser) readTagInt() (TagInt, error) {
	i, err := p.readInt32()
	return TagInt(i), err
}

func (p *Parser) readTagLong() (TagLong, error) {
	i, err := p.readInt64()
	return TagLong(i), err
}

func (p *Parser) readTagFloat() (TagFloat, error) {
	i, err := p.readFloat32()
	return TagFloat(i), err
}

func (p *Parser) readTagDouble() (TagDouble, error) {
	i, err := p.readFloat64()
	return TagDouble(i), err
}

func (p *Parser) readTagByteArray() (TagByteArray, error) {
	length, err := p.readInt32()
	if err != nil {
		return nil, err
	}

	bytes, err := p.readBytes(int(length))
	return bytes, err
}

func (p *Parser) readTagString() (TagString, error) {
	length, err := p.readUint16()
	if err != nil {
		return "", err
	}

	bytes, err := p.readBytes(int(length))
	return TagString(bytes), err
}

func (p *Parser) readTagList() (tag TagList, err error) {
	tag.ElementType, err = p.readTagId()
	if err != nil {
		return
	}

	var length int32
	length, err = p.readInt32()
	if err != nil {
		return
	}

	if length <= 0 {
		return
	}

	tag.Elements = make([]Tag, length)

	for i := int32(0); i < length; i++ {
		tag.Elements[i], err = tag.ElementType.Read(p)
		if err != nil {
			return
		}
	}

	return
}

func (p *Parser) readTagCompound() (TagCompound, error) {
	tag := make(TagCompound)

	for {
		// parse tag ID
		tagId, err := p.readTagId()
		if err != nil {
			return nil, err
		}

		if tagId == TagTypeEnd {
			break
		}

		// parse tag name
		var nameLength uint16
		nameLength, err = p.readUint16()
		if err != nil {
			return nil, err
		}

		var name string
		name, err = p.readUtf8(int(nameLength))

		subTag, err := tagId.Read(p)
		if err != nil {
			return nil, err
		}

		tag[name] = subTag
	}

	return tag, nil
}

func (p *Parser) readTagIntArray() (tag TagIntArray, err error) {
	var length int32
	length, err = p.readInt32()
	if err != nil {
		return
	}

	tag = make([]int32, length)

	for i := int32(0); i < length; i++ {
		tag[i], err = p.readInt32()
	}

	return
}

func (p *Parser) readTagLongArray() (tag TagLongArray, err error) {
	var length int32
	length, err = p.readInt32()
	if err != nil {
		return
	}

	tag = make([]int64, length)

	for i := int32(0); i < length; i++ {
		tag[i], err = p.readInt64()
	}

	return
}
