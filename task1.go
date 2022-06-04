//Есть коробка A * B * C
//Есть дверной проем X * Y
//Выяснить пройдет ли коробка через дверь

package main

import (
	"fmt"
)

func main() {
	var A, B, C int8
	fmt.Print("Enter box sizes (A, B, C): ")
	fmt.Scanf("%d %d %d", &A, &B, &C)
	fmt.Printf("Размеры коробки: %d %d %d", A, B, C)
}
