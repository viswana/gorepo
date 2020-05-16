package main

import (
	"fmt"

	"github.com/viswana/gorepo/dsalgo/sequences"
)

func main() {
	f := sequences.Fibonacci()
	for i := 0; i < 10; i += 1 {
		fmt.Printf("%v ", f())
	}
}
