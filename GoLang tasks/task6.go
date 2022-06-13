// Дан массив целых чисел из 200 элементов. Массив заполняется на основе датчика случайных чисел от 0 до 1000.
// Подсчитать количество четных и нечетных значений этого массива.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var odd, even int

func main() {
	arr :=make([]int, 200)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(arr); i++ {
		arr[i]=rand.Intn(1000)
	}
	for i, _ :=range arr {
		if arr[i]%2==0 {
			even++
		} else {
			odd++
		}
	}
fmt.Printf("В массиве %d чётных и %d нечётных значений", even, odd)
}