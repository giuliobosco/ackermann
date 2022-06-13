package main

import "fmt"

func ackerman(n, m int64) int64 {
	fmt.Println(n, m)
	if n == 0 {
		return m + 1
	}

	if m == 0 {
		return ackerman(n - 1, 1)
	}

	return ackerman(n - 1, ackerman(n, m - 1))
}

func main() {
	fmt.Println(ackerman(2, 1))
}