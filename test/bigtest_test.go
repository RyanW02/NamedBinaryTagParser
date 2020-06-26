package main

import (
	"bytes"
	"github.com/RyanW02/NamedBinaryTagParser/nbt"
	"os"
	"testing"
)

func TestBigTest(t *testing.T) {
	f, err := os.Open("bigtest.nbt")
	if err != nil {
		t.Errorf("failed to open bigtest.nbt: %x", err)
	}

	p := nbt.NewParser(f)
	compound, name, err := p.Read()
	if err != nil {
		t.Errorf("error parsing bigtest.nbt: %x", err)
	}

	if name != "Level" {
		t.Errorf("outer compound name doesn't match. got: %s wanted: Level", name)
	}

	// nested compound test
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "nested compound test",
		compound:  compound,
		tagType:   nbt.TagTypeCompound,
		tagValue: nbt.TagCompound{
			"egg": nbt.TagCompound{
				"name":  nbt.TagString("Eggbert"),
				"value": nbt.TagFloat(0.5),
			},
			"ham": nbt.TagCompound{
				"name":  nbt.TagString("Hampus"),
				"value": nbt.TagFloat(0.75),
			},
		},
	})

	// intTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "intTest",
		compound:  compound,
		tagType:   nbt.TagTypeInt,
		tagValue:  nbt.TagInt(2147483647),
	})

	// byteTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "byteTest",
		compound:  compound,
		tagType:   nbt.TagTypeByte,
		tagValue:  nbt.TagByte(127),
	})

	// stringTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "stringTest",
		compound:  compound,
		tagType:   nbt.TagTypeString,
		tagValue:  nbt.TagString("HELLO WORLD THIS IS A TEST STRING \xc3\x85\xc3\x84\xc3\x96!"),
	})

	// listTest (long)
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "listTest (long)",
		compound:  compound,
		tagType:   nbt.TagTypeList,
		tagValue: nbt.TagList{
			ElementType: nbt.TagTypeLong,
			Elements: []nbt.Tag{
				nbt.TagLong(11),
				nbt.TagLong(12),
				nbt.TagLong(13),
				nbt.TagLong(14),
				nbt.TagLong(15),
			},
		},
	})

	// doubleTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "doubleTest",
		compound:  compound,
		tagType:   nbt.TagTypeDouble,
		tagValue:  nbt.TagDouble(0.49312871321823148),
	})

	// floatTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "floatTest",
		compound:  compound,
		tagType:   nbt.TagTypeFloat,
		tagValue:  nbt.TagFloat(0.49823147058486938),
	})

	// longTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "longTest",
		compound:  compound,
		tagType:   nbt.TagTypeLong,
		tagValue:  nbt.TagLong(9223372036854775807),
	})

	// listTest (compound)
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "listTest (compound)",
		compound:  compound,
		tagType:   nbt.TagTypeList,
		tagValue: nbt.TagList{
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
	})

	// byteArrayTest
	byteArrayTestData := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		byteArrayTestData[i] = byte((i*i*255 + i*7) % 100)
	}

	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))",
		compound:  compound,
		tagType:   nbt.TagTypeByteArray,
		tagValue:  nbt.TagByteArray(byteArrayTestData),
	})

	// shortTest
	doFieldTest(fieldTestData{
		t:         t,
		fieldName: "shortTest",
		compound:  compound,
		tagType:   nbt.TagTypeShort,
		tagValue:  nbt.TagShort(32767),
	})
}

func BenchmarkBigTest(b *testing.B) {
	// open NBT file
	f, err := os.Open("bigtest.nbt")
	if err != nil {
		b.Errorf("failed to open bigtest.nbt: %x", err)
	}

	// read bytes
	data := make([]byte, 2048)
	n, err := f.Read(data)
	if err != nil {
		b.Errorf("failed to read bigtest.nbt: %x", err)
	}

	data = data[:n]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// stop timer & create reader
		b.StopTimer()
		reader := bytes.NewReader(data)
		b.StartTimer()

		p := nbt.NewParser(reader)

		// parse the data
		_, _, _ = p.Read()
	}
}
