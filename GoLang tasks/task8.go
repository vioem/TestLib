// Напишите программу с типами celsius, fahrenheit и kelvin и методами для конвертации из одного типа температуры в другой.
package main

import "fmt"

type celsius float64 
type  fahrenheit float64 
type kelvin float64

func KelTo(k kelvin) (celsius, fahrenheit) {
    return celsius(k - 273.15), fahrenheit((k - 273.15)* 9/5 + 32 )
}

func CelTo(c celsius) (fahrenheit, kelvin) {
    return fahrenheit((c - 32) * 5/9), kelvin(c + 273.15)
}

func FahrTo(f fahrenheit) (celsius, kelvin) {
    return celsius((f - 32) * 5/9), kelvin((f - 32) * 5/9 + 273.15)
}

var enter int

func main() {
	fmt.Println("What unit of measure will we be converting from?")
	fmt.Printf(" 1. celsius\n 2. fahrenheit\n 3. kelvin\n")
	fmt.Scanln(&enter)
	if enter==1 {
		var t celsius
		fmt.Println("Enter temperature in celsius:")
		fmt.Scanln(&t)
		fmt.Print(CelTo(t))
	} else if enter==2 {
		var t fahrenheit
		fmt.Println("Enter temperature in Fahrenheit:")
		fmt.Scanln(&t)
		fmt.Print(FahrTo(t))
	} else if enter==3 {
		var t kelvin
		fmt.Println("Enter temperature in Kelvin:")
		fmt.Scanln(&t)
		fmt.Print(KelTo(t))
	} else {
		fmt.Println("You entered an invalid value")
	} 
}
