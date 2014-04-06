package fib

import (
	"fmt"
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

func calc() <-chan string {

	c := make(chan string)

	go func() {

		c <- "Starting"

		res := fib(37)
		c <- fmt.Sprintf("Done: %v", res)

		close(c)

	}()

	return c
}

func p_fib(n int) {

	var c []<-chan string

	for i := 0; i < n; i++ {
		c = append(c, calc())
	}

	for {
		for i := 0; i < n; i++ {

			res, ok := <-c[i]
			if !ok {
				return
			}

			fmt.Println(res)
		}
	}

}
