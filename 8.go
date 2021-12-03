package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		n         int64
		nF        float64
		nDiff     float64
		match     int64
		matchF    float64
		SqrtN     float64
		SqrtClear float64
	)

	_, _ = fmt.Scan(&n)

	match = 1
	nF = float64(n)
	SqrtN = math.Sqrt(nF)
	SqrtClear = math.Pow(math.Trunc(SqrtN), 2)
	nDiff = nF - SqrtClear

	if n == 1 {
		fmt.Println(match * 4)
	} else if SqrtN-math.Trunc(SqrtN) < 1e-5 {
		matchF = float64(match)
		matchF = SqrtN * 2 * (SqrtN + 1)
		match = int64(matchF)
		fmt.Println(match)
	} else if n < 4 && n > 1 {
		match = match*4 + (n-1)*3
		fmt.Println(match)
	} else if nDiff == 1 && n > 4 {
		matchF = float64(match)
		matchF = nDiff*3 + math.Trunc(SqrtN)*2*(math.Trunc(SqrtN)+1)
		match = int64(matchF)
		fmt.Println(match)
	} else if nDiff > 1 && nDiff <= math.Trunc(SqrtN) && n > 4 {
		matchF = float64(match)
		matchF = math.Trunc(SqrtN)*2*(math.Trunc(SqrtN)+1) + 3 + (nDiff-1)*2
		match = int64(matchF)
		fmt.Println(match)
	} else if nDiff > math.Trunc(SqrtN) && n > 4 {
		matchF = float64(match)
		matchF = math.Trunc(SqrtN)*2*(math.Trunc(SqrtN)+1) + 6 + (nDiff-2)*2
		match = int64(matchF)
		fmt.Println(match)

	} else {

	}
}
