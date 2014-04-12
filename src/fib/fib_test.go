package fib

import (
	"fmt"
	"testing"
)

func TestFib1(t *testing.T) {

	f := FibR(37)
	f1 := Fib(37)

	if f != f1 {
		t.Error("Fib1 is not equal to Fib", f1, f)
	}
}

func BenchmarkFib1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(37)
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibR(37)
	}
}

func TestMain(t *testing.T) {

	for i := 0; i < 10; i++ {
		fmt.Printf("Result: %v\n", FibR(37))
	}

}

var tasks []int

func init_tasks(v int, n int) {

	tasks = tasks[0:0]

	for i := 0; i < n; i++ {
		tasks = append(tasks, v)
	}

}

func TestPFibR(t *testing.T) {
	init_tasks(37, 11)
	Runner(tasks, 10, FibR)
}

func TestPFib(t *testing.T) {
	init_tasks(37, 10)
	Runner(tasks, 10, Fib)
}
