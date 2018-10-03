package main

import (
	"bufio"
	"bytes"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/harikb/pghash/lib/pghash"
)

const buffLen = 100

func BenchmarkPghash(b *testing.B) {
	var buffers [][]byte
	r := rand.New(rand.NewSource(int64(os.Getpid())))
	for i := 0; i < b.N; i++ {
		buf := make([]byte, buffLen)
		n, err := r.Read(buf)
		if err != nil || n != buffLen {
			log.Fatalf("Can't create test dats %s", err)
		}
		buffers = append(buffers, buf)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		pghash.HashAny(buffers[i])
	}
}

func TestPghash(t *testing.T) {
	fp, err := os.Open("words_hash.txt")
	if err != nil {
		t.Errorf("Can't open words_hash.txt: %+v", err)
	}
	bfp := bufio.NewScanner(fp)
	lnum := 0
	for bfp.Scan() {
		lnum++
		line := bfp.Bytes()
		parts := bytes.Split(line, []byte("\t"))
		expectedInt, err := strconv.ParseInt(string(parts[1]), 10, 32)
		if err != nil {
			t.Errorf("bad data in words_hash.txt: %+v", err)
		}
		outputUnsigned, output := pghash.HashAny(parts[0])
		if output != int32(expectedInt) {
			t.Errorf("Error in line %d hash for key=>%s<, pg=%d, lib=%d (unsigned: %d)",
				lnum, parts[0], expectedInt, output, outputUnsigned)
			break
		}
	}
}
