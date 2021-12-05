package main

import (
	"fmt"
	"math"
)

func FullKube(exp int) int {
	var (
		result     int
		sideLen, _ = UpLow3(exp, "lower")
	)

	result = 3 * sideLen * int(math.Pow(float64(sideLen+1), 2))

	return result
}

func FullPlane(length int) int {
	var result int

	if length == 1 {
		result = 8
	} else {
		result = (length - 1) * (2*5 + (length-1)*3)
	}
	return result
}

func UpLow3(exp int, kind string) (int, int) {
	var f func(float64) float64
	switch kind {
	case "upper":
		f = math.Ceil
	case "lower":
		f = math.Floor
	default:
		f = math.Floor
	}
	//len of face of a cube                 amount of lil cube with face eq len
	return int(f(math.Cbrt(float64(exp)))), int(math.Pow(f(math.Cbrt(float64(exp))), 3))
}

func UpLow2(exp int, kind string) (int, int) {
	var f func(float64) float64
	switch kind {
	case "upper":
		f = math.Ceil
	case "lower":
		f = math.Floor
	default:
		f = math.Floor
	}
	//side length							plane area with side length eq len
	return int(f(math.Sqrt(float64(exp)))), int(math.Pow(f(math.Sqrt(float64(exp))), 2))
}

func LilPlane(exp int) int {
	var (
		_, lilArea       = UpLow2(exp, "lower")
		length, _    int = UpLow3(exp, "lower")
		upLen, _         = UpLow2(exp, "upper")
		_, area          = UpLow2(exp, "upper")
		_, totalArea     = UpLow2(length, "upper")
		result       int
	)

	if totalArea == area {
		result = FullPlane(upLen)
	}

	return result
}

func main() {
	var (
		n int
	)

	_, _ = fmt.Scan(&n)

	fmt.Println(FullKube(n))
}
