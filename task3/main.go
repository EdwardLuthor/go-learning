package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	times := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&times[i])
	}

	solved1 := 0
	time1 := 0
	for i := 0; i < n; i++ {
		if time1+times[i] <= 300 {
			time1 += times[i]
			solved1++
		} else {
			break
		}
	}

	solved2 := 0
	time2 := 0
	for i := n - 1; i >= 0; i-- {
		if time2+times[i] <= 300 {
			time2 += times[i]
			solved2++
		} else {
			break
		}

		if solved1 > solved2 {
			fmt.Println("Победил 1-ый студент")
		} else if solved2 > solved1 {
			fmt.Println("Победил 2-ой студент")
		} else {
			fmt.Println("Ничья")
		}

	}
}
