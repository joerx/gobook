package main

import (
	"testing"
)

func TestPopCountExp(t *testing.T) {
	tests := []struct {
		num uint64
		want uint8
	}{
		{2, 1},
		{16, 1},
		{25, 3},
		{412353451, 12},
	}
	
	for _, test := range tests {
		if res := PopCountExp(test.num); res != test.want {
			t.Errorf("PopCount(%d) = %d, want %d", test.num, res, test.want)
		}
	}
}

const testVal uint64 = 43592343

func BenchmarkPopCountExp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountExp(testVal)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(testVal)
	}
}
