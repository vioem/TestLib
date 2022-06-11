package main

import (
	"fmt"
)

pivot := rand.Int() % len(a)

func main() {
	ar := []int{3, 4, 1, 2, 5, 7, -1, 0}
	Quicksort(ar)
	fmt.Println(ar)
}


func Quicksort(ar []int) {
	if len(ar) <= 1 {
		return
	}

	split := partition(ar)

	Quicksort(ar[:split])
	Quicksort(ar[split:])
}

func partition(ar []int) int {

	//  x x x x x x x p X X X X X X X X
	//                ^                  <- n
	//  p x x x x x x x X X X X X X X X
	//  s p x x x ...                    <- n
	//
	// n^2

	pivot := ar[len(ar)/2]

	left := 0
	right := len(ar) - 1

	for {
		for ; ar[left] < pivot; left++ {
		}

		for ; ar[right] > pivot; right-- {
		}

		if left >= right {
			return right
		}

		swap(ar, left, right)
	}
}

func swap(ar []int, i, j int) {
	tmp := ar[i]
	ar[i] = ar[j]
	ar[j] = tmp
}


func mysqlQuery(query string, args ...any) *sql.Result {
	defer tracer.StartSpan("mysqlRequest").Finish()
	return db.Query(query, args...)
}
