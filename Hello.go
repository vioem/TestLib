package main

import (
	"fmt"
)

func main() {
	var room = "озеро"

	switch room { // Выражения для каждого случая
	case "пещера":
		fmt.Println("Вы находитесь в тускло освещенной пещере.")
	case "озеро":
		fmt.Println("Лед кажется достаточно крепким.")
		fallthrough // Переходит на следующий случай
	case "глубина":
		fmt.Println("Вода такая холодная, что сводит кости.")
	}

	var i interface{}
	i = room
	switch i := i.(type) {
	case string:
		fmt.Printf("string: %s", i)
	case int:
		fmt.Printf("sdfsdf")
	default:
		fmt.Printf("something else: %T:%v", i, i)
	}
}

/*
   CMP A, B
   JMP -> 5
   CMP V, B
   JMP -> 7
   CMP W, B
   JMP -> 9
   ...
5:   MOV
6:   CALL
7:   MOV
8:   CALL
9:   MOV
     JMP -> END
A:   CALL
B:   MOV
C:   CALL
END: sdfsdf


*/
