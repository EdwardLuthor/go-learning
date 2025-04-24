package main

import "fmt"

func main() {
	var N, A, B, S int
	_, err := fmt.Scan(&N, &A, &B)
	if err != nil {
		return
	}
	if N >= 100 || A >= 100 || B >= 100 {
		fmt.Println("ERROR: Ошибка ввода, все значения должна быть меньше либо равны 100 ")
		return
	}
	S = A * B * N * 2
	fmt.Println(S)
}
