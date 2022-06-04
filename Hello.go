package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var cnt1 = 0
	var cnt2 = 0

	var arr [200]int

	for i := range arr {
		arr[i] = rand.Intn(1000)

	}

	for i := range arr {
		if arr[i]%2 == 0 {
			cnt1++
		} else {
			cnt2++
		}
	}
	fmt.Println(cnt1, cnt2)
}
