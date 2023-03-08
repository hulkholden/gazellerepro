package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	ints := []int{4, 3, 2, 1}
	slices.Sort(ints)
	fmt.Printf("Ints: %v\n", ints)
}
