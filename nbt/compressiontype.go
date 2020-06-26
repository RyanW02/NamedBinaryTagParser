package nbt

type CompressionType int

const (
	Uncompressed CompressionType = iota
	Gzip
	Zlib
)
