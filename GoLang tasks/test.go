package main

import (
	"fmt"
)

func main() {
	ar := []int{1, 2, 3}
	ar2 := []int{2, 3, 4}
	ar3 := Intersection(ar, ar2)
	fmt.Println(ar3, find(1, ar))
}

func Intersection(val1, val2 []int) []int {
	x := len(val1)
	val3 := []int{}
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if val1[i] == val2[j] && find(i, val3) != i {
				val3 = append(val3, val1[i])
			}
		}
	}
	return val3
}

func find(enter int, value []int) (yes int) {
	yes = 0
	for i := range value {
		if value[i] == enter {
			yes++
		} else {
			continue
		}
	}
	return yes
}
