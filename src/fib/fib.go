//fib package allows to calculate nth fibonacci number.
package fib

//it is recursive, that's why it is  very slow
func FibR(n int) int {

	if n == 0 {
		return 0
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

	if n == 0 {
		return a
	}
	if n == 1 {
		return b
	}

	for i := 2; i <= n; i++ {

		res = a + b

		a = b
		b = res

	}

	return res
}

//tasks are being done by given number of workers
//channels are mailboxes here - async
func Runner(tasks []int, workers int, f func(int) int) []int {

	in, out := make(chan int, 10), make(chan int, 10)

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

	return res

}
