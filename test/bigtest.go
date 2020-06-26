package main

import "github.com/RyanW02/NamedBinaryTagParser/nbt"

var bigTest = nbt.TagCompound{
	"nested compound test": nbt.TagCompound{
		"egg": nbt.TagCompound{
			"name":  nbt.TagString("Eggbert"),
			"value": nbt.TagFloat(0.5),
		},
		"ham": nbt.TagCompound{
			"name":  nbt.TagString("Hampus"),
			"value": nbt.TagFloat(0.75),
		},
	},
	"intTest":    nbt.TagInt(2147483647),
	"byteTest":   nbt.TagByte(127),
	"stringTest": nbt.TagString("HELLO WORLD THIS IS A TEST STRING \xc3\x85\xc3\x84\xc3\x96!"),
	"listTest (long)": nbt.TagList{
		ElementType: nbt.TagTypeLong,
		Elements: []nbt.Tag{
			nbt.TagLong(11),
			nbt.TagLong(12),
			nbt.TagLong(13),
			nbt.TagLong(14),
			nbt.TagLong(15),
		},
	},
	"doubleTest": nbt.TagDouble(0.49312871321823148),
	"floatTest":  nbt.TagFloat(0.49823147058486938),
	"longTest":   nbt.TagLong(9223372036854775807),
	"listTest (compound)": nbt.TagList{
		ElementType: nbt.TagTypeCompound,
		Elements: []nbt.Tag{
			nbt.TagCompound{
				"created-on": nbt.TagLong(1264099775885),
				"name":       nbt.TagString("Compound tag #0"),
			},
			nbt.TagCompound{
				"created-on": nbt.TagLong(1264099775885),
				"name":       nbt.TagString("Compound tag #1"),
			},
		},
	},
	"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))": nbt.TagByteArray(byteArrayTestData()),
	"shortTest": nbt.TagShort(32767),
}

func byteArrayTestData() []byte {
	byteArrayTestData := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		byteArrayTestData[i] = byte((i*i*255 + i*7) % 100)
	}
	return byteArrayTestData
}
