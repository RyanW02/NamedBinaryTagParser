package main

import (
	"github.com/RyanW02/NamedBinaryTagParser/nbt"
	"reflect"
	"testing"
)

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
		t.Errorf("%v value doesn't match. got: %v wanted: %v", name, value, wanted)
	}
}
