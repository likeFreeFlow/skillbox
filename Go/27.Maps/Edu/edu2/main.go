package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{
		1: "first",
		2: "second",
		3: "third",
		4: "forth",
	}
	for k := range m {
		fmt.Println(k)
	}
	sliceOfKeys := make([]int, 0)
	for k := range m {
		sliceOfKeys = append(sliceOfKeys, k)
	}
	sortedSliceOfKeys := sort.IntSlice(sliceOfKeys)

	for _, v := range sortedSliceOfKeys {
		fmt.Println(v, m[v])
	}
}
