// Дан массив некоторых целых чисел, найдите значение 20 в нем и, если оно присутствует, заменить его на 200.
// Обновите список только при первом вхождении числа 20.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 30)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(arr); i++ {
		arr[i] = rand.Intn(len(arr))
	}

	for i, _ := range arr {
		if arr[i]==20 {
			arr[i]=200
			break
		}
	}
	fmt.Println(arr)
}