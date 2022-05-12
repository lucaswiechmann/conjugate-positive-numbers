package main

import (
	"flag"
	"fmt"
	"strings"

	"conjugate-positive-numbers/cmd/algorithm"
)

func main() {
	fmt.Println("Conjugate Positive Numbers")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Example: ./conjugate 6-4-2")
		return
	}

	arrNumbers := strings.Split(flag.Args()[0], "-")

	conjugateArrayNumbers := algorithm.ConjugateNumbers(arrNumbers)
	fmt.Println(conjugateArrayNumbers)
}
