// Поменять местами самый большой и самый маленький элементы массива.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr :=make([]int, 10)
	var min,max, minn, maxn int
	rand.Seed(time.Now().UnixNano())
	for i:=0; i<len(arr); i++ {
		arr[i]=rand.Intn(len(arr))
	}
	for i, _ :=range arr {
		if i==0 || arr[i] < min {
			min=arr[i]
		}
		if i==0 || arr[i] > max {
			max=arr[i]
		}
	}
	for i, _ :=range arr {
		if arr[i]==max {
			maxn=i
		}
		if arr[i]==min {
			minn=i
		}
	}
	fmt.Println("Массив до: ", arr, "\nМаксимум: ", max, "\nМинимум: ", min)
	arr[minn], arr[maxn] = arr[maxn], arr[minn]
	fmt.Println("Массив после: ", arr)
}