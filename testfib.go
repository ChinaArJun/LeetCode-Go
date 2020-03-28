package test

import (
	"fmt"
	"time"
)

func fib(f int) int {
	if f <= 2 {
		return f
	}
	return fib(f-1) + fib(f-2)
}

func test() {
	t := time.Now()
	fmt.Println("hello")
	fib(30)
	elapsed := time.Since(t)
	fmt.Println("app run time", elapsed)
}
