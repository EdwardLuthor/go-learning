package main

import "fmt"

func romanToInt(s string) int {
	r := 0
	i := 0

	if s[i] == 'I' {
		r += 1
	} else if s[i] == 'V' {
		r += 5
	} else if s[i] == 'X' {
		r += 10
	} else if s[i] == 'L' {
		r += 50
	} else if s[i] == 'C' {
		r += 100
	} else if s[i] == 'D' {
		r += 500
	} else if s[i] == 'M' {
		r += 1000
	}
	i++
	return r
}

func main() {
	var romanNumeral string
	_, err := fmt.Scanln(&romanNumeral)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	integerValue := romanToInt(romanNumeral)
	fmt.Println("Целое число:", integerValue)
}
