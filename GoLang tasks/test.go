package main

import "fmt"

func main() {
	coralSlice := []string{"blue coral", "foliose coral", "pillar coral", "elkhorn coral"}

	coralSlice = append(coralSlice[:1], coralSlice[4:]...)

	fmt.Printf("%q\n", coralSlice)
}
