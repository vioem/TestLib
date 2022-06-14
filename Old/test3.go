
package main

import (
    "fmt"
)

func main() {
    var temperatures = []float64{
        -28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
    }
    set := make(map[float64]bool) // Создание карты с булевыми значениями
    for _, t := range temperatures {
        set[t] = true
    }
    if set[-28.0] {
        fmt.Println("часть множества") // Выводит: часть множества
    }
    fmt.Println(set) // Выводит: Prints map[-31:true -29:true -23:true -33:true -28:true 32:true]
     
}