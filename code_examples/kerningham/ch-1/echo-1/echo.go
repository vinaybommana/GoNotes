package main

import (
	"fmt"
	"os"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
	int32AsIntValues := make([]int, len(a))
	for i, val := range a {
		int32AsIntValues[i] = int(val)
	}
	sort.Ints(int32AsIntValues)

	// maxLength
	return 1
}

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
	pickingNumbers([]int32{4, 6, 5, 3, 3, 1})
}
