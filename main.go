package main

import (
	"fmt"
)

func main() {
	k := 0
	arr := [5]int{2, 13, 16, 7, 9}
	for i := 0; i < 5; i++ {

		if arr[i]%2 == 0 {
			k += arr[i]
		}
	}
	fmt.Println(k)

}
