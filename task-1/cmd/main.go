package main

import (
	"dersnek/fibonacci/pkg/fibonacci"
	"flag"
	"fmt"
)

var nFlag = flag.Int("n", 0, "index of a desired Fibonacci number. Min: 0, max: 20.")

func main() {
	flag.Parse()

	if *nFlag < 0 || *nFlag > 20 {
		fmt.Println("Index outside of accepted range. Must be in range: 0..20.")
		return
	}

	fmt.Printf("Fib(%d) = %d\n", *nFlag, fibonacci.Calculate(*nFlag))
}
