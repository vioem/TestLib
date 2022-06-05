//Есть коробка A * B * C
//Есть дверной проем X * Y
//Выяснить пройдет ли коробка через дверь?
package main

import (
	"fmt"
)

func main() {
	var A, B, C, x, y int
	var aa, bb, cc bool
	fmt.Println("Введите параметры коробки:")
	fmt.Scanf("%d %d %d\n", &A, &B, &C)
	fmt.Println("Введите параметры дверного проёма:")
	fmt.Scanf("%d %d\n", &x, &y)
	aa = A <= x || A <= y
	bb = B <= x || B <= y
	cc = C <= x || C <= y
	if aa && bb && cc {
		fmt.Println("Пройдет")
	} else {
		fmt.Println("не пройдет")
	}
}
