package main

import (
	"fmt"
	"math"
)

func Race(v1, v2, g int) [3]int {
	var (
		result [3]int
		t      float64
	)

	if v1 < v2 {
		t = float64(g) / float64(v2-v1)
		for i := 0; i < 3; i++ {
			result[i] = int(math.Floor(t*math.Pow(60, float64(i)))) % 60
		}

	} else {
		result = [3]int{-1, -1, -1}
	}

	return result
}

func main() {
	fmt.Println(Race(720, 850, 70))
	fmt.Println(Race(820, 81, 550))
	fmt.Println(Race(80, 91, 37))
}
