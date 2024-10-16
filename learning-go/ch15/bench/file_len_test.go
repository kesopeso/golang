package bench_test

import (
	"ch15/bench"
	"fmt"
	"math/rand"
	"os"
	"testing"
)

const fileLenFile = "testdata/filelen.txt"

func TestMain(m *testing.M) {
	makeData()
	exitValue := m.Run()
	removeData()
	os.Exit(exitValue)
}

var blackhole int

func BenchmarkFileLen(b *testing.B) {
	bufferSizes := []int{1, 10, 100, 1000, 10000, 100000}
	for _, bs := range bufferSizes {
		b.Run(fmt.Sprintf("FileLen%d", bs), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := bench.FileLen(fileLenFile, bs)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}

func TestFileLen(t *testing.T) {
	len, err := bench.FileLen(fileLenFile, 5)
	if err != nil {
		t.Fatal(err)
	}
	if len != 65204 {
		t.Errorf("wrong file len read, expected: 22, got: %d", len)
	}
}

func removeData() {
	os.Remove(fileLenFile)
}

func makeData() {
	file, err := os.Create(fileLenFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	repeatableRand := rand.New(rand.NewSource(1))
	for i := 0; i < 10000; i++ {
		data := makeWord(repeatableRand, repeatableRand.Intn(10)+1)
		file.Write(data)
	}
}

func makeWord(r *rand.Rand, l int) []byte {
	out := make([]byte, l+1)
	for i := 0; i < l; i++ {
		out[i] = 'a' + byte(r.Intn(26))
	}
	out[l] = '\n'
	return out
}
