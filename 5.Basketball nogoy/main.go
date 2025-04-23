package main

import "fmt"

func main() {
	totalScore1 := 0
	totalScore2 := 0

	for i := 0; i < 4; i++ {
		var a, b int
		fmt.Scanln(&a, &b)

		if a < 0 || a > 100 || b < 0 || b > 100 {
			fmt.Println("ERROR: Некорректный ввод очков (должно быть от 0 до 100).")
			return
		}

		totalScore1 += a
		totalScore2 += b
	}

	if totalScore1 > totalScore2 {
		fmt.Println(1)
	} else if totalScore2 > totalScore1 {
		fmt.Println(2)
	} else {
		fmt.Println("DRAW")
	}
}
