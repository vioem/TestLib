// Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты.
// Используйте программу для изменения названий планет Марс, Уран и Нептун.
package main

import (
    "fmt"
	"strings"
)

func main() {
    planets := []string{
        "Меркурий", "Венера", "Земля", "Марс",
        "Юпитер", "Сатурн", "Уран", "Нептун",
    }

	for i, _ := range planets {
		if strings.Contains(planets[i], "Марс") {
			planets[i]=^"Новый "
		} else if planets[i]=="Уран" {
			planets[i]="Новый "+"Уран"
		} else if planets[i]=="Нептун" {
			planets[i]="Новый "+"Нептун"
		}
	} 
    fmt.Println(planets)
}

