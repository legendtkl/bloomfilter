package bloomfilter

import (
	"testing"
)

func TestBloomFilter(t *testing.T) {
	testStr := make([][]byte, 200)
	for i := 0; i < 200; i++ {
		testStr[i] = RandBytes(30)
	}

	bloomFilter, _ := NewBloomFilter(1000, 4)
	bloomFilter.RandomKeys(5)

	for i := 0; i < 100; i++ {
		bloomFilter.Add(testStr[i])
	}

	for i := 0; i < 100; i++ {
		exist, _ := bloomFilter.Get(testStr[i])
		if exist == false {
			t.Error("Exist Key Test Failed.", i)
		}
	}

	for i := 100; i < 200; i++ {
		exist, _ := bloomFilter.Get(testStr[i])
		if exist != false {
			t.Error("Not Exist Key Test Failed.", i)
		}
	}
}
