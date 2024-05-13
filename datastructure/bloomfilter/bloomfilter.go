package bloomfilter

import "hash/crc32"

func hash(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}

type BloomFilter struct {
	hash func([]byte) uint32
	cap  int
}

func New(cap int) *BloomFilter {
	bf := &BloomFilter{
		hash: hash,
		cap:  cap,
	}
	return bf
}
