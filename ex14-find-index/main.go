package main

import "fmt"

func strStr(haystack string, needle string) int {

	haystackLen := len(haystack)
	needleLen := len(needle)

	if needleLen > haystackLen {
		return -1
	}

	for i := 0; i <= haystackLen-needleLen; i++ {
		match := true
		for j := 0; j < needleLen; j++ {

			if haystack[i+j] != needle[j] {
				match = false
				break
			}

		}
		if match {
			if i == 0 {
				return 0
			} else {
				return -1
			}
		}
	}
	return -1
}

func main() {
	var haystack string
	fmt.Printf("Введите строку haystack: ")
	fmt.Scanln(&haystack)

	var needle string
	fmt.Printf("Введите строку needle: ")
	fmt.Scanln(&needle)

	result := strStr(haystack, needle)
	fmt.Println(result)
}
