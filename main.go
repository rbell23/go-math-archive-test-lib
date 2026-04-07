package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rbell23/go-math-archive-test-lib/mathlib"
)

func main() {
	a := flag.Int("a", 2, "first number")
	b := flag.Int("b", 3, "second number")
	flag.Parse()

	if flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "unexpected positional arguments")
		os.Exit(2)
	}

	fmt.Println(mathlib.Add(*a, *b))
}
