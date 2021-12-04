package main

import (
	"fmt"
	"math"
)

func ClearSqr3(exp int) int {
	var (
		expFloat = float64(exp)
		ResFloat float64
	)

	ResFloat = math.Cbrt(expFloat)
	ResFloat = math.Floor(ResFloat)

	return int(ResFloat)
}

func FullKube(exp int) int {
	var (
		result  int
		SideLen = ClearSqr3(exp)
	)

	result = 3 * SideLen * int(math.Pow(float64(SideLen+1), 2))

	return result
}

func UpLow(exp int, kind string) (int, int) {
	var f func(float64) float64
	switch kind {
	case "upper":
		f = math.Ceil
	case "lower":
		f = math.Floor
	default:
		f = math.Floor
	}
	//len of face of a cube                 amount of lil cube inside a big one with face eq len
	return int(f(math.Cbrt(float64(exp)))), int(f(math.Pow(math.Cbrt(float64(exp)), 3)))
}

func main() {
	var (
		n int
	)

	_, _ = fmt.Scan(&n)

	fmt.Println(FullKube(n))
}
