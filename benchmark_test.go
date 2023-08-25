package main

import (
	"os"
	"testing"
)

func benchmarkMergeBySort(file string, b *testing.B) {
	f, err := os.Open(file)
	defer f.Close()

	if err != nil {
		b.Fatalf("Failed to open test file %s: %v", file, err)
	}

	data, err := ReadIntervals(f)
	if err != nil {
		b.Fatalf("Failed to read intervals from file %s: %v", file, err)
	}

	b.ResetTimer() // reset the timer after reading from the file

	b.Run(file, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			MergeBySort(data)
		}
	})
}

func BenchmarkMergeBySortLowXS(b *testing.B) {
	benchmarkMergeBySort("./testdata/low.xs.txt", b)
}

func BenchmarkMergeBySortLowS(b *testing.B) {
	benchmarkMergeBySort("./testdata/low.s.txt", b)
}

func BenchmarkMergeBySortLowXL(b *testing.B) {
	benchmarkMergeBySort("./testdata/low.l.txt", b)
}

func BenchmarkMergeBySortMidM(b *testing.B) {
	benchmarkMergeBySort("./testdata/mid.m.txt", b)
}

func BenchmarkMergeBySortMidL(b *testing.B) {
	benchmarkMergeBySort("./testdata/mid.l.txt", b)
}

func BenchmarkMergeBySortHighXL(b *testing.B) {
	benchmarkMergeBySort("./testdata/high.xl.txt", b)
}
