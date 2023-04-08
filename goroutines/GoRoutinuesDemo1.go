package main

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}


func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x - 1) + fib(x - 2)
}


func main() {
	go spinner(100 * time.Microsecond)
	const n = 10
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)


}