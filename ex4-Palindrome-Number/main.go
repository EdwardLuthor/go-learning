package main

import "fmt"

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	originalX := x
	reversedX := 0
	for x > 0 {
		lastDigit := x % 10
		reversedX = reversedX*10 + lastDigit
		x /= 10
	}
	return originalX == reversedX
}

func main() {
	var number int
	fmt.Scanln(&number)
	if isPalindrome(number) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
