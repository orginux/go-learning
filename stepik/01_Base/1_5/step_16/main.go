package main

import "fmt"

const (
	hDegree = 360 / 12        // 30 degree per hr
	mDegree = (12 * 60) / 360 // 2 minutes is 1 degree
)

func main() {
	var degree int
	fmt.Scan(&degree)
	hours := degree / hDegree
	minutes := mDegree * (degree % hDegree)
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
