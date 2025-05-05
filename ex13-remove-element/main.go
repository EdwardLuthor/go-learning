package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for _, v := range nums {
		if v != val {
			nums[k] = v
			k++
		}
	}
	return k
}

func main() {

	var n int
	fmt.Printf("Введите количество элементов:")
	fmt.Scanln(&n)
	numbers := make([]int, n)
	fmt.Printf("Введите %d элементов через пробел: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	var dummy string
	fmt.Scanln(&dummy)

	var value int
	fmt.Printf("Введите значение которое требуется убрать из массива:")
	fmt.Scanln(&value)

	k := removeElement(numbers, value)
	fmt.Println("Количество элементов не равных", value, ":", k)
	fmt.Println("Массив после удаления:", numbers[:k])
}
