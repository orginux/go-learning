package main

import "fmt"

func main() {
	var i int
	fmt.Scan(&i)
	switch {
	case i < 0:
		fmt.Println("Число отрицательное")
	case i > 0:
		fmt.Println("Число положительное")
	default:
		fmt.Println("Ноль")
	}
}
