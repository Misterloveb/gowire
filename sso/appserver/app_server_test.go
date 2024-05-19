package appserver

import (
	"fmt"
	"math"
	"testing"
	"unsafe"
)

func TestApp1(t *testing.T) {
	app := NewApp("app1", "app1_123456", "www.app1.com:8081")
	app.Run()
}

func TestApp2(t *testing.T) {
	app := NewApp("app2", "app2_123456", "www.app2.com:8082")
	app.Run()
}
func Change(b []byte) string {
	fmt.Printf("b-- %x", &b)
	return *(*string)(unsafe.Pointer(&b))
}

type BloomFilter struct {
	bitArray []bool
	numHash  uint
}

func NewBloomFilter(capacity uint, falsePositiveRate float64) *BloomFilter {
	bitSize := math.Ceil((-float64(capacity) * math.Log(falsePositiveRate)) / math.Pow(math.Log(2), 2))
	numHash := uint(math.Ceil((bitSize / float64(capacity)) * math.Log(2)))

	return &BloomFilter{
		bitArray: make([]bool, int(bitSize)),
		numHash:  numHash,
	}
}
func TestChanl(t *testing.T) {
	_ = NewBloomFilter(1024*1024, 0.01)
}
