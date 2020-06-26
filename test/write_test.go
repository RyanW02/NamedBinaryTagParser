package main

import (
	"bytes"
	"github.com/RyanW02/NamedBinaryTagParser/nbt"
	"testing"
)

func TestWrite(t *testing.T) {
	outerName := "Level"

	// encode data
	var buf bytes.Buffer
	writer := nbt.NewWriter(&buf)
	if err := writer.Write(bigTest, outerName); err != nil {
		t.Fatalf("failed encoding data: %v", err)
	}

	// decode data
	parser := nbt.NewParser(&buf)
	decoded, decodedName, err := parser.Read()
	if err != nil {
		t.Fatalf("failed deconding test data: %v", err)
	}

	if decodedName != outerName {
		t.Fatalf("outer names do not match. got: %s wanted: %s", decodedName, outerName)
	}

	valueCheck(t, "bigtest", decoded, bigTest)
}

func BenchmarkWrite(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// stop timer & create writer
		b.StopTimer()

		var buf bytes.Buffer
		writer := nbt.NewWriter(&buf)

		b.StartTimer()

		// write the data
		_ = writer.Write(bigTest, "Level")
	}
}
