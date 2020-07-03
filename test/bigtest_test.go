package main

import (
	"bytes"
	"github.com/RyanW02/NamedBinaryTagParser/nbt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	f, err := os.Open("bigtest.nbt")
	if err != nil {
		t.Errorf("failed to open bigtest.nbt: %v", err)
	}

	p, err := nbt.NewParser(f)
	if err != nil {
		t.Errorf("error creating parser: %v", err)
	}

	compound, name, err := p.Read()
	if err != nil {
		t.Errorf("error parsing bigtest.nbt: %v", err)
	}

	if name != "Level" {
		t.Errorf("outer compound name doesn't match. got: %s wanted: Level", name)
	}

	byteArrayTestData := make([]byte, 1000)
	for i := 0; i < 1000; i++ {
		byteArrayTestData[i] = byte((i*i*255 + i*7) % 100)
	}

	valueCheck(t, "bigtest", compound, bigTest)
}

func BenchmarkParse(b *testing.B) {
	// open NBT file
	f, err := os.Open("bigtest.nbt")
	if err != nil {
		b.Errorf("failed to open bigtest.nbt: %v", err)
	}

	// read bytes
	data := make([]byte, 2048)
	n, err := f.Read(data)
	if err != nil {
		b.Errorf("failed to read bigtest.nbt: %v", err)
	}

	data = data[:n]

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// stop timer & create reader
		b.StopTimer()

		reader := bytes.NewReader(data)
		p, err := nbt.NewParser(reader)
		if err != nil {
			b.Errorf("error creating parser: %v", err)
		}

		b.StartTimer()

		// parse the data
		_, _, _ = p.Read()
	}
}
