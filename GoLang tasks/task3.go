// Написать программу: дан массив из 10 чисел, которые задаются с помощью датчика случайных чисел.
// Программа находит повторяющиеся элементы и удалеяет их их массива. Например дан массив {1,1,1,2,3,4,2,3,4} результат {1,2,3,4}.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 10)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(arr); i++ {
		arr [i]=rand.Intn(10)
	}
	fmt.Println("Массив: ", arr, "\nРезультат: ", uniq(arr))
}

func uniq(val []int) []int {
	m := make(map[int]struct{})
	var m2 []int

	for _, item := range val {
		m[item] = struct{}{}
	}

	for item :=range m {
		m2 =append(m2, item)
	}
	return m2
}
