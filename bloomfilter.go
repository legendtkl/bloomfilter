package bloomfilter

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"

	"github/legendtkl/bitmap"
)

type BloomFilter struct {
	bitMap   *bitmap.BitMap
	m        uint64
	k        uint64
	keys     [][]byte
	hashFunc func(message, key []byte) uint64
}

// NewBloomFilter create a new bloom filter
func NewBloomFilter(size, hashN uint64) (*BloomFilter, error) {
	bitMap, _ := bitmap.NewBitMap(size)
	filter := &BloomFilter{
		bitMap:   bitMap,
		m:        size,
		k:        hashN,
		hashFunc: hmacHash,
	}

	return filter, nil
}

// InitHashFunc init the hash func
func (filter *BloomFilter) InitHashFunc(f func(message, key []byte) uint64) {
	filter.hashFunc = f
}

// InitKeys init hash Keys
func (filter *BloomFilter) InitKeys(hashKeys [][]byte) {
	filter.keys = hashKeys
}

// RandomKeys generate hash keys randomly
func (filter *BloomFilter) RandomKeys(keyLength int) {
	var hashKeys [][]byte
	for i := 0; i < int(filter.k); i++ {
		randBytes := RandBytes(10)
		hashKeys = append(hashKeys, randBytes)
	}
	filter.keys = hashKeys
}

// Add add an item
func (filter *BloomFilter) Add(msg []byte) error {
	for _, v := range filter.keys {
		val := filter.hashFunc(msg, v)
		filter.bitMap.SetOne(val % filter.m)
	}
	return nil
}

// Get get the item, if exist, return true; else, return false
func (filter *BloomFilter) Get(msg []byte) (bool, error) {
	for _, v := range filter.keys {
		val := filter.hashFunc(msg, v)
		if x, _ := filter.bitMap.GetPosition(val % filter.m); x != 1 {
			return false, nil
		}
	}
	return true, nil
}

// hmacHash is default hash function
func hmacHash(msg, key []byte) uint64 {
	mac := hmac.New(sha256.New, key)
	mac.Write(msg)
	res := binary.BigEndian.Uint64(mac.Sum(nil))
	return res
}

// RandBytes generate random bytes of length (length)
func RandBytes(length int) []byte {
	data := make([]byte, length)
	rand.Read(data)

	return data
}
