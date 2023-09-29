package main

import (
	"fmt"
)

func main() {
	k := 0
	arr := [5]int{}
	for i := 0; i < 5; i++ {
		fmt.Scan(arr[i])
		if arr[i]%2 == 0 {
			k += arr[i]
		}
	}
	fmt.Println(k)

}
