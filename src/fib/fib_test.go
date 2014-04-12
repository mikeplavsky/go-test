package fib

import (
	"fmt"
	"testing"
)

func ExampleFib() {

	fmt.Println(
		Fib(0),
		Fib(1),
		Fib(2),
		Fib(3),
		Fib(4),
		Fib(5))
	// Output: 0 1 1 2 3 5

}

func ExampleFibR() {

	fmt.Println(
		FibR(0),
		FibR(1),
		FibR(2),
		FibR(3),
		FibR(4),
		FibR(5))
	// Output: 0 1 1 2 3 5

}

func TestFib1(t *testing.T) {

	f := FibR(37)
	f1 := Fib(37)

	if f != f1 {
		t.Error("Fib1 is not equal to Fib", f1, f)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(37)
	}
}

func BenchmarkFibR(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibR(37)
	}
}

func test_perf(t *testing.T, f func(int) int, nth int, times int) {

	tasks := make([]int, 0)

	for i := 0; i < times; i++ {
		tasks = append(tasks, nth)
	}

	res := Runner(tasks, 10, f)

	if len(res) != times {
		t.Error("wrong slice length")
	}

	if res[0] != f(nth) {
		t.Error("wrong fibonacci number")
	}

}

func TestPFibR(t *testing.T) {
	test_perf(t, FibR, 38, 11)
}

func TestPFib(t *testing.T) {
	test_perf(t, Fib, 38, 11)
}
