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

func calc(c chan int) {
  c <- fib(37)
}

func p_fib(n int) {

	c := make(chan int)

        for i := 0; i < n; i++ {
          go calc(c)
        }

        res := 0

        for i := 0; i < n; i++ {
          res = <-c
        }

        fmt.Println("Done!\n", res)
}
