package nbt

type Tag interface {
	Type() TagId
}

type TagEnd struct{}
type TagByte byte
type TagShort int16
type TagInt int32
type TagLong int64
type TagFloat float32
type TagDouble float64
type TagByteArray []byte
type TagString string
type TagList struct {
	ElementType TagId
	Elements    []Tag
}
type TagCompound map[string]Tag
type TagIntArray []int32
type TagLongArray []int64

func (t TagEnd) Type() TagId {
	return TagTypeEnd
}

func (t TagByte) Type() TagId {
	return TagTypeByte
}

func (t TagShort) Type() TagId {
	return TagTypeShort
}

func (t TagInt) Type() TagId {
	return TagTypeInt
}

func (t TagLong) Type() TagId {
	return TagTypeLong
}

func (t TagFloat) Type() TagId {
	return TagTypeFloat
}

func (t TagDouble) Type() TagId {
	return TagTypeDouble
}

func (t TagByteArray) Type() TagId {
	return TagTypeByteArray
}

func (t TagString) Type() TagId {
	return TagTypeString
}

func (t TagList) Type() TagId {
	return TagTypeList
}

func (t TagCompound) Type() TagId {
	return TagTypeCompound
}

func (t TagIntArray) Type() TagId {
	return TagTypeIntArray
}

func (t TagLongArray) Type() TagId {
	return TagTypeLongArray
}
