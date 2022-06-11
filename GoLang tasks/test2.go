package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 4, 1, 2, 5, 7, -1, 0} // len 8
	BubbleSort2(arr)
	fmt.Println(arr)
}

func BubbleSort2(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := len(arr) - 1; j > i; j-- {
			if arr[j-1] > arr[j] {
				swap(arr, j-1, j)
			}
		}
	}
}

func swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
