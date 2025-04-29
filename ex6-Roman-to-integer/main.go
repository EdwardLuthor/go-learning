package main

import "fmt"

func romanToIntString(s string) int {
	romanValues := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	result := 0
	for i := 0; i < len(s); i++ {
		currentSymbol := string(s[i])
		currentValue := romanValues[currentSymbol]
		if i+1 < len(s) {
			nextSymbol := string(s[i+1])
			nextValue := romanValues[nextSymbol]
			if currentValue < nextValue {
				result -= currentValue
			} else {
				result += currentValue
			}
		} else {
			result += currentValue
		}
	}
	return result
}

func main() {
	var romanNumeral string
	_, err := fmt.Scanln(&romanNumeral)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	integerValue := romanToIntString(romanNumeral)
	fmt.Println("Целое число:", integerValue)
}
