package main

import "fmt"

func removeDuplicates(nums []int) []int {

	result := make([]int, 0)
	result = append(result, nums[0])

	for i := 1; i < len(nums); i++ {
		prev := nums[i-1]
		current := nums[i]
		if prev != current {
			result = append(result, current)
		}

	}
	return result

}

func main() {
	var n int
	fmt.Printf("Введите количество элементов:")
	fmt.Scanln(&n)

	nums := make([]int, n)
	fmt.Printf("Введите %d элементов через пробел: ", n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	result := removeDuplicates(nums)
	fmt.Println("Массив с удаленными дубликатами:", result)
	fmt.Println("Количество уникальных элементов:", len(result))
}
