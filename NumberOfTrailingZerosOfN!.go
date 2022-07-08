package main

import (
	"fmt"
	"math"
)

func Zeros(n int) int {
	count := 0

	if n < 5 {
		fmt.Println("Error")
	}

	for i := 1; i <= 20; i++ {
		count = count + int(math.Trunc(float64(n)/math.Pow(float64(5), float64(i))))
	}

	return count
}

func main() {
	fmt.Println(Zeros(0))
	fmt.Println(Zeros(6))
	fmt.Println(Zeros(30))
	fmt.Println(Zeros(678))
	fmt.Println(Zeros(4568))
	fmt.Println(Zeros(132473))
}
