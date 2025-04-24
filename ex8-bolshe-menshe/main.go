package main

import "fmt"

func main() {
	var X, Y int
	_, err := fmt.Scan(&X, &Y)
	if err != nil {
		return
	}
	if X > Y {
		fmt.Println(">")
	} else if X < Y {
		fmt.Println("<")
	} else {
		fmt.Println("=")
	}
}
