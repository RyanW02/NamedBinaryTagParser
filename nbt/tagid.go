package nbt

type TagId uint8

const (
	TagTypeEnd TagId = iota
	TagTypeByte
	TagTypeShort
	TagTypeInt
	TagTypeLong
	TagTypeFloat
	TagTypeDouble
	TagTypeByteArray
	TagTypeString
	TagTypeList
	TagTypeCompound
	TagTypeIntArray
	TagTypeLongArray
)

func (id TagId) Read(parser *Parser) (tag Tag, err error) {
	switch id {
	case TagTypeEnd:
		tag = TagEnd{}
	case TagTypeByte:
		tag, err = parser.readTagByte()
	case TagTypeShort:
		tag, err = parser.readTagShort()
	case TagTypeInt:
		tag, err = parser.readTagInt()
	case TagTypeLong:
		tag, err = parser.readTagLong()
	case TagTypeFloat:
		tag, err = parser.readTagFloat()
	case TagTypeDouble:
		tag, err = parser.readTagDouble()
	case TagTypeByteArray:
		tag, err = parser.readTagByteArray()
	case TagTypeString:
		tag, err = parser.readTagString()
	case TagTypeList:
		tag, err = parser.readTagList()
	case TagTypeCompound:
		tag, err = parser.readTagCompound()
	case TagTypeIntArray:
		tag, err = parser.readTagIntArray()
	case TagTypeLongArray:
		tag, err = parser.readTagLongArray()
	}

	return
}
