package main

import (
	"fmt"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func sumOfEven(s []int) int {
	k := 0
	for i := 0; i < 5; i++ {
		if s[i]%2 == 0 {
			k += s[i]
		}
	}
	return k
}

func main() {
	max := maxOccur("❤❤aabccc❤❤cdd😊❤❤")
	fmt.Println(max)
}

func maxOccur(s string) int {
	counts := make(map[string]int)
	for _, v := range s {
		counts[string(v)]++
	}
	return slices.Max(maps.Values(counts))
}
