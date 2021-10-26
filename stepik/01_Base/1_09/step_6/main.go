package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := n / 100
	b := n / 10 % 10
	c := n % 10
	if a != b && b != c && c != a {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")
}
