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

func p_fib(tasks []int, workers int, f func(int) int) {

	in, out := make(chan int), make(chan int)

	for i := 0; i < workers; i++ {
		go func() {
			for {
				out <- f(<-in)
			}
		}()
	}

	for _, v := range tasks {
		in <- v
	}

	var res []int

	for i := 0; i < len(tasks); i++ {
		res = append(res, <-out)
	}

	fmt.Println("Done!\n", res)
}
