# NamedBinaryTagParser  
As the name suggets, NamedBinaryTagParser is an NBT parser (& now also a serializer), written in Go.

It supports all tag types, as well as being able to intelligently detect and handle gzip & zlib compression
(or no compression).

# Getting Started
## Tag IDs
NamedBinaryTagParser's tag ID constants can be found
[here](https://github.com/RyanW02/NamedBinaryTagParser/blob/master/nbt/tagid.go).

## Tag types
NamedBinaryTagParser's tag types can be found
[here](https://github.com/RyanW02/NamedBinaryTagParser/blob/master/nbt/tags.go).

## Parsing
### Simple
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

### More Complex Example (list of compounds)
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

### Json Example
A small example in which we serialize an NBT compound to a JSON object.
```go
func example(r io.Reader) {
	parser := nbt.NewParser(r)
	tag, name, err := parser.Read()
	if err != nil {
		panic(err)
	}

	withName := map[string]interface{}{
		name: tag,
	}

	marshalled, err := json.Marshal(withName)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshalled))
}
```

<details>
    <summary>JSON output</summary>
    
    ````javascript
    {
    	"Level": {
    		"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))": "",
    		"byteTest": 127,
    		"doubleTest": 0.4931287132182315,
    		"floatTest": 0.49823147,
    		"intTest": 2147483647,
    		"listTest (compound)": {
    			"ElementType": 10,
    			"Elements": [{
    				"created-on": 1264099775885,
    				"name": "Compound tag #0"
    			}, {
    				"created-on": 1264099775885,
    				"name": "Compound tag #1"
    			}]
    		},
    		"listTest (long)": {
    			"ElementType": 4,
    			"Elements": [11, 12, 13, 14, 15]
    		},
    		"longTest": 9223372036854775807,
    		"nested compound test": {
    			"egg": {
    				"name": "Eggbert",
    				"value": 0.5
    			},
    			"ham": {
    				"name": "Hampus",
    				"value": 0.75
    			}
    		},
    		"shortTest": 32767,
    		"stringTest": "HELLO WORLD THIS IS A TEST STRING ÅÄÖ!"
    	}
    }
    ````
</details>

## Serialization
When serializing data, if you wish to use compression, you should parse a [gzip writer](https://golang.org/pkg/compress/gzip/#Writer)
or [zlib writer](https://golang.org/pkg/compress/zlib/#Writer) to NamedBinaryTagParser, unlike when parsing data, in which
NamedBinaryTagParser automatically detects the compression and applies the correct reader.

### Example
```go
func example(w io.Writer) {
	// your data set
	// this is wiki.vg's bigtest.nbt
	data := nbt.TagCompound{
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
		"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))": nbt.TagByteArray(byteArrayTestData),
		"shortTest": nbt.TagShort(32767),
	}

	// create out writer
	writer := nbt.NewWriter(w)
	outerName := "Level" // this is the name of the outer NBT compound
	if err := writer.Write(data, outerName); err != nil {
		panic(err)
	}
}
```
  
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
BenchmarkParse-12          53491             21447 ns/op
BenchmarkWrite-12         204349              5781 ns/op
PASS
ok      github.com/RyanW02/NamedBinaryTagParser/test    8.542s
```

# Todo

 - More tests
