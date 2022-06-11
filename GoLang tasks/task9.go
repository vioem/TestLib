/* Написать программу работы с статическими массивами
Даны два статических массива  - одинаковый длины, например [1,2,3], [4,5,6]. Реализовать в виде функций следующие операции:
- инициализация статических массивов осуществляется на основе датчика случайных чисел
- поиск заданного числа в массиве, [1,1,2]  - число 1
- подсчет количества четных и нечетных чисел
- поиск min, max в массиве
- сортировка значений массива по возрастанию
- поиск пересечений двух массивов, [1,2,3], [2,3,4] - - >> [2,3]
- сложение значений двух массивов
Замечание 1. Использовать глобальные переменные  - можно!
Замечание 2. Показать передачу массивов в виде аргумента для функции, показать возвращение результата работы функции в виде массива
Замечание 3.При реализации сортировки запрещено использовать встроенные функции в GO
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

var enter, odd, even, min, max, x int

func main() {
	// инициализация статических массивов осуществляется на основе датчика случайных чисел
	arr1, arr2 := generator(10)

	// поиск заданного числа в массиве, например [1,1,2]  - число 1
	fmt.Println("Какое число хотите найти в данных массивах?:")
	fmt.Scanln(&enter)
	fmt.Printf("Число %d повторяется в первом массиве %d раз, во втором %d раз\n", enter, find(enter, arr1), find(enter, arr2))

	// подсчет количества четных и нечетных чисел
	even, odd = num(arr1)
	fmt.Printf("Количество четных чисел в массиве %d, нечетных %d\n", even, odd)
	even, odd = num(arr2)
	fmt.Printf("Количество четных чисел в массиве %d, нечетных %d\n", even, odd)

	// поиск min, max в массиве
	fmt.Printf("Минимальное значение %d. Максимальное значение %d\n", ext(arr1)[0], ext(arr1)[1])
	fmt.Printf("Минимальное значение %d. Максимальное значение %d\n", ext(arr2)[0], ext(arr2)[1])

	// сортировка значений массива по возрастанию
	fmt.Println("Отсортированные массивы", BubbleSort(arr1), BubbleSort(arr2))

	// поиск пересечений двух массивов
	fmt.Println("В двух массивах есть следующие элементы: ", Intersection(arr1, arr2))

	// сложение значений двух массивов,
	fmt.Println("Сумма двух массивов: ", Plus(arr1, arr2))
}

func generator(x int) ([]int, []int) {
	val1 := make([]int, x, x)
	val2 := make([]int, x, x)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < x; i++ {
		val1[i] = rand.Intn(x)
		val2[i] = rand.Intn(x)
	}
	return val1, val2
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

func num(value []int) (odd int, even int) {
	even = 0
	odd = 0
	for i := range value {
		if value[i]%2 == 0 {
			even++
		} else {
			odd++
			continue
		}
	}
	return even, odd
}

func ext(value []int) []int {
	for i, j := range value {
		if i == 0 || j < min {
			min = j
		}
	}
	for i, j := range value {
		if i == 0 || j > max {
			max = j
		}
	}
	return []int{min, max}
}

func BubbleSort(value []int) []int {
	for i := 0; i < len(value); i++ {
		for j := i + 1; j < len(value); j++ {
			if value[i] > value[j] {
				value[i], value[j] = value[j], value[i]
			}
		}
	}
	return value
}

func Intersection(val1, val2 []int) []int {
	x := len(val1)
	val3 := []int{}
	for i := 0; i < x; i++ {
		for j := 0; j < x; j++ {
			if val1[i] == val2[j] /* && find(i, val3) != i  */ {
				val3 = append(val3, val1[i])
			}
		}
	}
	return val3
}

func Plus(val1, val2 []int) []int {
	val3 := make([]int, len(val1), len(val1))
	for i := 0; i < len(val1); i++ {
		val3[i] = val1[i] + val2[i]
	}
	return val3
}
