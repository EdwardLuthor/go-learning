package main

import "fmt"

func main() {

	var N, M, K int
	_, err := fmt.Scan(&N, &M, &K)
	if N >= 100 || M >= 100 || K >= 10000 {
		fmt.Println("ERROR: Ошибка ввода, N,M значения не должны быть больше 100, K не должно быть больше 10000 ")
		return
	}
	if err != nil {
		return
	}
	totalOreshki := N * M
	if totalOreshki >= K {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")

	}

}
