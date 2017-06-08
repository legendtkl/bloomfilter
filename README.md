# bloomfilter
go implemented bloom filter.
about bloom filter, please refer to https://en.wikipedia.org/wiki/Bloom_filter

# example
```go
func main() {
    bloomFilter, _ := bloomfilter.NewBloomFilter(1000, 4)
    bloomFilter.RandomKeys(10)

    url := []string{"http://legendtkl.com/2017/04/28/golang-gc/",
        "http://legendtkl.com/2017/04/02/golang-alloc/",
        "http://legendtkl.com/2017/03/21/malloc-os-knowledge/",
        "http://legendtkl.com/2016/12/31/git-good-practice-gitflow/"}
    bloomFilter.Insert([]byte(url[0]))
    bloomFilter.Insert([]byte(url[1]))
    bloomFilter.Insert([]byte(url[2]))

    res, _ := bloomFilter.Lookup([]byte(url[3]))
    fmt.Println(res)
}
```

# notice
Because bloom filter is not 100% correct, in other word sometimes will return exist for an not-exist key. So when you run `go test`, if the test fail, that is reasonable.