package fib

import (
	"fmt"
	"time"
)

func fib(n int) int {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}

func fib1(n int) int {

	res := 0

	a := 0
	b := 1

	for i := 0; i < n; i++ {

		res = a + b

		a = b
		b = res

	}

	return res
}

func main() {

	fmt.Printf("Starting...\n")

	start := time.Now()

	for i := 0; i < 10; i++ {
		fmt.Printf("Result: %v\n", fib(37))
	}

	elapsed := time.Since(start)
	fmt.Printf("Elapsed %s\n", elapsed)

}
