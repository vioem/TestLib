
package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
    newWorlds := make([]string, len(worlds)) // Создает новый срез вместо прямого изменения worlds

    for i := range worlds {
        newWorlds[i] = prefix + " " + worlds[i]
    }
    return newWorlds
}

func main() {
	planets := []string{"Венера", "Марс", "Юпитер"}
	newPlanets := terraform("Нью", planets...)
	fmt.Println(newPlanets) // Выводит: [Нью Венера Нью Марс Нью Юпитер]	
}