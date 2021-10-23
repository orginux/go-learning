package main

import "fmt"

func main() {
	var n, max, count int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		if n >= max {
			if n != max {
				count = 0
			}
			max = n
			count++
		}

	}
	fmt.Println(count)
}
