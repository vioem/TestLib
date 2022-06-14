// Напишите программу для преобразования слайса строки через добавление слова "Новый " перед названием планеты.
// Используйте программу для изменения названий планет Марс, Уран и Нептун.
package main

import (
    "fmt"
	"strings"
)

func main() {
    planets := []string{ "Меркурий", "Венера", "Земля", "Марс", "Юпитер", "Сатурн", "Уран", "Нептун"}
    fmt.Println(terr("Новый", planets...))
}

func terr(prefix string, worlds ...string) []string {
    newWorlds := make([]string, len(worlds))

    for i := range worlds {
        if strings.Contains(worlds[i], "Марс") || strings.Contains(worlds[i], "Нептун") || strings.Contains(worlds[i], "Уран") {
			newWorlds[i] = prefix + " " + worlds[i]
		} else {
			newWorlds[i] = worlds[i]
		}
    }
    return newWorlds
}

