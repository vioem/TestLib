// Написать программу: дан массив из 10 чисел, которые задаются с помощью датчика случайных чисел.
// Программа находит повторяющиеся элементы и считает количество повторений. Например дан массив {1,1,1,2,3,4,2,3,4} результат число 1,
// повторяется 3 раза, число 2 два раза, числа 3 два раза, число 4 два раза.
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	arr := [10]int{}
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		arr[i]=rand.Intn(10)
	}
	
	var sum int
	m := map[int]int{}

	for _, x := range arr {
		for z:=0; z < len(arr); z++ {
			if x == arr[z] {
				sum++
			}
		}
		m[x] = sum
		sum = 0
	}
	
	/* fmt.Println(arr)
	fmt.Println(m) */
	for i,j := range m {
		fmt.Printf("Число %d повторяется %d раз(а)\n", i, j)
	}
}