package main

import "testing"

func BenchmarkCalcHashAndCopyFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calcHashAndCopyFile()
	}
}
