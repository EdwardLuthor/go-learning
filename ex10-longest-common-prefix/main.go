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
	var word1, word2, word3, word4, word5 string

	fmt.Printf("Введите первое слово: ")
	fmt.Scanln(&word1)

	fmt.Printf("Введите второе слово: (или нажмите Enter, если нет): ")
	fmt.Scanln(&word2)

	fmt.Printf("Введите третье слово: (или нажмите Enter, если нет): ")
	fmt.Scanln(&word3)

	fmt.Printf("Введите четвертое слово: (или нажмите Enter, если нет) ")
	fmt.Scanln(&word4)

	fmt.Printf("Введите пятое слово: (или нажмите Enter, если нет) ")
	fmt.Scanln(&word5)

	var strs []string
	if word1 != "" {
		strs = append(strs, word1)

	}
	if word2 != "" {
		strs = append(strs, word2)

	}
	if word3 != "" {
		strs = append(strs, word3)

	}
	if word4 != "" {
		strs = append(strs, word4)

	}
	if word5 != "" {
		strs = append(strs, word5)

	}

	result := PrefixFinder(strs)
	fmt.Printf("Саммый длинный общий префикс: %s\n", result)
}
