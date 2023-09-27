package main

import (
	"fmt"
)

func main() {
	k := 0
	arr := [5]int{1, 3, 6, 12, 5}
	for i := 0; i < 5; i++ {
		if arr[i]%2 == 0 {
			k += arr[i]
		}
	}
	fmt.Println(k)

}
