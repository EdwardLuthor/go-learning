package main

import "fmt"

func PrefixFinder(strs []string) string {
	firstWord := strs[0]
	firstWordLength := len(firstWord)
	CommonPrefix := ""
	for i := 0; i < firstWordLength; i++ {
		var currentLetter rune = rune(firstWord[i])
		for j, otherWord := range strs {
			if j == 0 {
				continue
			}
			if len(otherWord) <= i || rune(otherWord[i]) != currentLetter {
				return CommonPrefix
			}
		}
		CommonPrefix += string(currentLetter)
	}
	return CommonPrefix
}

func main() {
	var numWords int
	fmt.Printf("Введите количество слов: ")
	_, err := fmt.Scanln(&numWords)
	if err != nil {
		fmt.Println("Ошибка ввода количество слов: ", err)
		return
	}

	if numWords <= 0 {
		fmt.Println("Пожалуйста, введите положительное количество слов. ")
		return
	}

	strs := make([]string, numWords)
	fmt.Println("Введите слова по одному:")
	for i := 0; i < numWords; i++ {
		fmt.Printf("Слово %d: ", i+1)
		_, err := fmt.Scanln(&strs[i])
		if err != nil {
			fmt.Println("Ошибка ввода слова:", err)
			return
		}
	}
	result := PrefixFinder(strs)
	fmt.Printf("Саммый длинный общий префикс: %s", result)
}
