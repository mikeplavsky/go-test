package fib

import "fmt"

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

func p_fib(n int, f func(int) int) {

	in, out := make(chan int), make(chan int)

	for i := 0; i < n; i++ {
		go func() {
			for i := range in {
				out <- f(i)
			}
		}()
	}

	for i := 0; i < n; i++ {
		in <- 37
	}

	var res []int

	for i := 0; i < n; i++ {
		res = append(res, <-out)
	}

	fmt.Println("Done!\n", res)
}
