//fib package allows to calculate nth fibonacci number.
package fib

import "fmt"

//it is recursive, that's why it is  very slow
func FibR(n int) int {

	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}

	return FibR(n-1) + FibR(n-2)
}

//it is fast: based on for
func Fib(n int) int {

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

//tasks are being done by given number of workers
func Runner(tasks []int, workers int, f func(int) int) {

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
