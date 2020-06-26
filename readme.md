# NamedBinaryTagParser  
As the name suggets, NamedBinaryTagParser is an NBT parser, written in Go.

It supports all tag types, as well as being able to intelligently detect and handle gzip & zlib compression
(or no compression).

# Getting Started  
## Simple
```go
func example(reader io.Reader) {
	// create our parser instance
	// takes an io.Reader
	parser := nbt.NewParser(reader)

	// the Parser.Read() function will completely parse the NBT compound provided from the reader
	// it returns the parsed NBT compound, which is a type alias of map[string]nbt.Tag,
	// the name of the outer NBT compound, and also the error if one occurred.
	compound, name, err := parser.Read()
	if err != nil {
		panic(err)
	}

	// say in our compound we have a Long tag called "my_long"
	myTag := compound["my_long"] // this returns the generic tag

	// we can verify that the tag is a long
	if myTag.Type() != nbt.TagTypeLong {
		panic("tag is not a long")
	}

	// we can now safely cast to a nbt.TagLong, which is a type alias for int64
	myLong := myTag.(nbt.TagLong)
	fmt.Println(myLong)
}
```

## More Complex Example (list of compounds)
```go
func example(reader io.Reader) {
	// create our parser instance
	parser := nbt.NewParser(reader)

	// read the NBT compound
	compound, name, err := parser.Read()
	if err != nil {
		panic(err)
	}

	// say this time we have a list of compounds, where the list is called "my_list"
	myTag := compound["my_list"] // this returns the generic tag

	// we can verify that the tag is a list
	if myTag.Type() != nbt.TagTypeList {
		panic("tag is not a list")
	}

	// cast to nbt.TagList
	myList := myTag.(nbt.TagList)

	// verify that the elements are compounds
	if myList.ElementType != nbt.TagTypeCompound {
		panic("elements are not compounds")
	}

	// myList.Elements is a []nbt.Tag
	for _, tag := range myList.Elements {
		// we have verified that they are compounds already, so we can cast without further checks
		compound := tag.(nbt.TagCompound)

		// we can treat the compound like normal now
		// for example, i'll retrieve a string called "my_string"
		myString, ok := compound["my_string"].(nbt.TagString)
		if !ok {
			panic("my_string was not a string")
		}

		fmt.Println(myString)
	}
}
```

## Tag IDs
NamedBinaryTagParser's tag ID constants can be found
[here](https://github.com/RyanW02/NamedBinaryTagParser/blob/master/nbt/tagid.go).

## Tag types
NamedBinaryTagParser's tag types can be found
[here](https://github.com/RyanW02/NamedBinaryTagParser/blob/master/nbt/tags.go).
  
# Testing  
I have written a test based on [wiki.vg's bigtest.nbt](https://wiki.vg/NBT#bigtest.nbt). You can run the test with
```bash
go test ./test
```

# Benchmarking
I have also written a benchmark based on [wiki.vg's bigtest.nbt](https://wiki.vg/NBT#bigtest.nbt). You can run the benchmark with
```bash
go test ./test -bench .
```

## Results
Running the benchmark on my home PC (R5 3600X, 12GB RAM), I achieved the following results:
```bash
$ go test ./test -bench .
goos: linux
goarch: amd64
pkg: github.com/RyanW02/NamedBinaryTagParser/test
BenchmarkBigTest-12        55897             21273 ns/op
PASS
ok      github.com/RyanW02/NamedBinaryTagParser/test    2.839s
```

# Todo

 - More tests
 - Creating NBT data, as well as parsing it
