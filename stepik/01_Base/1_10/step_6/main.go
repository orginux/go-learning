package main

import "fmt"

func main() {
	var n, m, sum int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&m)
		if m > 9 && m < 100 && m%8 == 0 {
			sum += m
		}
	}
	fmt.Println(sum)
}
