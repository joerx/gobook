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
			t.Errorf("PopCount(%b) = %d, want %d", test.num, res, test.want)
		}
	}
}

func TestPopCountShift(t *testing.T) {
	var i uint64
	for ; i < 100; i++ {
		if got, want := PopCountShift(i), PopCountExp(i); got != want {
			t.Errorf("PopCount(%b) = %d, want %d", i, got, want)
		}
	}
}

func TestPopCountFoo(t *testing.T) {
	var i uint64
	for ; i < 1000; i++ {
		if got, want := PopCountFoo(i), PopCountExp(i); got != want {
			t.Errorf("PopCountFoo(%b) = %d, want %d", i, got, want)
		}
	}
}

const testVal uint64 = 3428352343634235

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

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(testVal)
	}
}

func BenchmarkPopCountFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountFoo(testVal)
	}
}
