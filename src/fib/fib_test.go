package fib

import "testing"

func TestFib1(t *testing.T) {

	f := fib(37)
	f1 := fib1(37)

	if f != f1 {
		t.Error("Fib1 is not equal to Fib", f1, f)
	}
}

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib1(37)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(37)
	}
}
