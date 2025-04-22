package main

import (
	"fmt"
)

func isValidSquare(s string) bool {
	if len(s) != 2 {
		return false
	}
	col := s[0]
	row := s[1]
	return col >= 'A' && col <= 'H' && row >= '1' && row <= '8'
}

func isKnightMove(move string) string {
	if len(move) != 5 || move[2] != '-' {
		return "ERROR"
	}

	start := move[:2]
	end := move[3:]
	if !isValidSquare(start) || !isValidSquare(end) {
		return "ERROR"
	}
	colDiff := abs(int(end[0]) - int(start[0]))
	rowDiff := abs(int(end[1]) - int(start[1]))
	if colDiff == 2 && rowDiff == 1 || colDiff == 1 && rowDiff == 2 {
		return "Yes"
	}
	return "No"
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	var userMove string
	//	log.Printf("%d", int('H')-int('A')) //
	fmt.Scanln(&userMove)
	result := isKnightMove(userMove)
	fmt.Println(result)
}
