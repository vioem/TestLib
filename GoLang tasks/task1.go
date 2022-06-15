//Есть коробка A * B * C
//Есть дверной проем X * Y
//Выяснить пройдет ли коробка через дверь?
package main

import	"fmt"

func main() {
	var a,b,c,d,f int
	fmt.Println("Enter box a*b*c sizes:")
	fmt.Scanf("%d %d %d\n", &a, &b, &c)
	fmt.Println("Eter door x*y size:")
	fmt.Scanf("%d %d\n", &d, &f)
	arr :=[]int{a,b,c}
	enter(arr,d,f)
}
 
func enter(val []int, x,y int) {
	smap := make(map[int][]int) 
	for _, j :=range  val{
		if j <=x {	
			smap[x] = append(smap[x], j)
		} else if j <= y {
			smap[y] = append(smap[y], j)
		}
	}
	if smap[x] != nil && smap[y] != nil {
		fmt.Println("Pass")
	} else {
		fmt.Println("Not pass")

	}
}