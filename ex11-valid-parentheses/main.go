package main

import "fmt"

func isValid(s string) bool {

	circleBracket := []int{}
	squareBracket := []int{}
	figureBracket := []int{}

	for i := 0; i < len(s); i++ {
		bracket := s[i]
		if bracket == '(' {
			circleBracket = append(circleBracket, i)
		} else if bracket == '[' {
			squareBracket = append(squareBracket, i)
		} else if bracket == '{' {
			figureBracket = append(figureBracket, i)
		} else if bracket == ')' {
			if len(circleBracket) > 0 {
				circleBracket = circleBracket[:len(circleBracket)-1]
			} else {
				return false
			}
		} else if bracket == ']' {
			if len(squareBracket) > 0 {
				squareBracket = squareBracket[:len(squareBracket)-1]
			} else {
				return false
			}
		} else if bracket == '}' {
			if len(figureBracket) > 0 {
				figureBracket = figureBracket[:len(figureBracket)-1]
			} else {
				return false
			}
		}
	}
	return len(circleBracket) == 0 && len(squareBracket) == 0 && len(figureBracket) == 0
}

func main() {
	var numTests int
	fmt.Print("Введите количество строк для проверки: ")
	fmt.Scanln(&numTests)

	for i := 0; i < numTests; i++ {
		var input string
		fmt.Printf("Введите строку скобок %d: ", i+1)
		fmt.Scanln(&input)

		result := isValid(input)
		fmt.Printf("isValid(\"%s\") = %t\n", input, result)
	}
}
