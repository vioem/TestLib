// Написать программу: дан массив из 10 чисел, которые задаются с помощью датчика случайных чисел.
// Программа находит повторяющиеся элементы и удалеяет их их массива. Например дан массив {1,1,1,2,3,4,2,3,4} результат {1,2,3,4}.

// Golang Program that Removes Duplicate
// Elements From the Array

package main

import "fmt"

func uniq(arr []int) []int {
	occurred := map[int]bool{}
	result := []int{}
	for e := range arr {

		// check if already the mapped
		// variable is set to true or not
		if occurred[arr[e]] != true {
			occurred[arr[e]] = true

			// Append to result slice.
			result = append(result, arr[e])
		}
	}

	return result
}

func main() {
	array1 := []int{1, 5, 3, 4, 1, 6, 6, 6, 8, 7, 13, 5}
	fmt.Println(array1)
	unique_items := unique(array1)
	fmt.Println(unique_items)
}
