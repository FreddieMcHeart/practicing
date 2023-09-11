package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n     int64
		nDiff float64
		sqrtN float64
	)

	_, _ = fmt.Scan(&n)

	sqrtN = math.Sqrt(float64(n))
	nDiff = float64(n) - math.Pow(math.Trunc(sqrtN), 2)

	switch {
	case n == 1:
		fmt.Println(4)
	case sqrtN-math.Trunc(sqrtN) < 1e-5:
		fmt.Println(int64(sqrtN * 2 * (sqrtN + 1)))
	case n < 4 && n > 1:
		fmt.Println(4 + (n-1)*3)
	case nDiff == 1 && n > 4:
		fmt.Println(int64(nDiff*3 + math.Trunc(sqrtN)*2*(math.Trunc(sqrtN)+1)))
	case nDiff > 1 && nDiff <= math.Trunc(sqrtN) && n > 4:
		fmt.Println(int64(math.Trunc(sqrtN)*2*(math.Trunc(sqrtN)+1) + 3 + (nDiff-1)*2))
	case nDiff > math.Trunc(sqrtN) && n > 4:
		fmt.Println(int64(math.Trunc(sqrtN)*2*(math.Trunc(sqrtN)+1) + 6 + (nDiff-2)*2))
	}
}
