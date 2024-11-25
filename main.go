package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"github.com/bits-and-blooms/bloom/v3"
)

func main() {
	filter := bloom.NewWithEstimates(10_000_000, 0.01)
	filter.Add([]byte("Love"))
	filter.Add([]byte("Death"))
	fmt.Println(filter.Test([]byte("Love")))

	var n uint = 10_000_000
	fmt.Println(bloom.EstimateFalsePositiveRate(20*n, 5, n) > 0.001)

	expectedFpRate := 0.01
	m, k := bloom.EstimateParameters(n, expectedFpRate)
	actualFpRate := bloom.EstimateFalsePositiveRate(m, k, n)
	fmt.Printf("expectedFpRate=%v, actualFpRate=%v\n", expectedFpRate, actualFpRate)

	b, _ := filter.MarshalBinary()
	fw, _ := os.Create("filter")
	fw.Write(b)
	fw.Close()

	fr, _ := os.Open("filter")
	var buf bytes.Buffer
	buf.ReadFrom(bufio.NewReader(fr))

	var filter2 bloom.BloomFilter // := bloom.NewWithEstimates(10_000_000, 0.01)
	filter2.UnmarshalBinary(buf.Bytes())
	fr.Close()

	fmt.Println(filter2.Test([]byte("Love")))
	fmt.Println(filter2.Test([]byte("Death")))

	filter3 := bloom.NewWithEstimates(10_000_000, 0.001)
	fw, _ = os.Create("filter3")
	bin, _ := filter3.MarshalBinary()
	fw.Write(bin)
	fw.Close()
}
