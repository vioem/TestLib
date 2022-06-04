// Golang program to illustrate the usage of
// fmt.Scanf() function

// Including the main package
package main

// Importing fmt
import (
	"fmt"
)

// Calling main
func main() {

	// Declaring some variables
	var name string
	var alphabet_count int

	// Calling Scanf() function for
	// scanning and reading the input
	// texts given in standard input
	fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &alphabet_count)

	// Printing the given texts
	fmt.Printf("The word %s containing %d number of alphabets.",
		name, alphabet_count)

}
