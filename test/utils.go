package main

import (
	"github.com/RyanW02/NamedBinaryTagParser/nbt"
	"reflect"
	"testing"
)

type fieldTestData struct {
	t         *testing.T
	fieldName string
	compound  nbt.TagCompound
	tagType   nbt.TagId
	tagValue  nbt.Tag
}

func doFieldTest(data fieldTestData) {
	tag := data.compound[data.fieldName]
	typeCheck(data.t, data.fieldName, tag, data.tagType)
	valueCheck(data.t, data.fieldName, tag, data.tagValue)
}

func typeCheck(t *testing.T, name string, tag nbt.Tag, wanted nbt.TagId) {
	if tag.Type() != wanted {
		t.Errorf("%s type doesn't match. got: %d wanted: %d", name, tag.Type(), wanted)
	}
}

func valueCheck(t *testing.T, name string, value interface{}, wanted interface{}) {
	var passed bool
	if _, isCompound := wanted.(nbt.TagCompound); isCompound {
		passed = reflect.DeepEqual(value, wanted)
	} else if _, isList := wanted.(nbt.TagList); isList {
		passed = reflect.DeepEqual(value, wanted)
	} else if _, isByteArray := wanted.(nbt.TagByteArray); isByteArray {
		passed = reflect.DeepEqual(value, wanted)
	} else {
		passed = value == wanted
	}

	if !passed {
		t.Errorf("%s value doesn't match. got: %d wanted: %d", name, value, wanted)
	}
}
