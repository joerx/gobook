package main

import (
	"testing"

	"github.com/joerx/gobook/exercises/2.3/popcount1"
	"github.com/joerx/gobook/exercises/2.3/popcount2"
)

var tests = []struct {
	num  uint64
	want int
}{
	{2, 1},
	{16, 1},
	{25, 3},
	{412353451, 12},
}

const benchVal uint64 = 3428352343634235

func TestPopCount1(t *testing.T) {
	for _, v := range tests {
		if got := popcount1.PopCount(v.num); got != v.want {
			t.Errorf("Wanted %d, got %d", v.want, got)
		}
	}
}

func TestPopCount2(t *testing.T) {
	for _, v := range tests {
		if got := popcount2.PopCount(v.num); got != v.want {
			t.Errorf("Wanted %d, got %d", v.want, got)
		}
	}
}

func BenchmarkPopCount1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount1.PopCount(benchVal)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount2.PopCount(benchVal)
	}
}
