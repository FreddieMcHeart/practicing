package main

import (
	"fmt"
	"math"
)

// FullKube calculate total matches from cube
func FullKube(exp int) int {
	var (
		result     int
		sideLen, _ = UpLow3(exp, "lower")
	)

	result = 3 * sideLen * int(math.Pow(float64(sideLen+1), 2))

	return result
}

// FullPlane calculate total matches from square
func FullPlane(length int) int {
	return (length-1)*(2*5+(length-1)*3) + 8
}

// UpLow3 calculate side len of bigger and smaller cube and its volume
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

// UpLow2 calculate side len of bigger and smaller square and its area
func UpLow2(length int, kind string) (int, int) {
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
	return int(f(math.Sqrt(float64(length)))), int(math.Pow(f(math.Sqrt(float64(length))), 2))
}

// LilPlane func takes total lilCubes on area and len of face of a Big Cube
func LilPlane(exp int, upLen int) int {
	var (
		lilLen, lilArea = UpLow2(exp, "lower")
		length, area    = UpLow2(exp, "upper")
		totalArea       = int(math.Pow(float64(upLen), 2))
		diff            = area - (area - exp + lilArea)
		result          int
	)
	//
	//fmt.Println("lilLen ", lilLen)
	//fmt.Println("lilArea ", lilArea)
	//fmt.Println("length ", length)
	//fmt.Println("area ", area)
	//fmt.Println("totalArea ", totalArea)
	//fmt.Println("diff ", diff)

	if totalArea == exp {
		result = FullPlane(length)
	} else if lilLen == length {
		result = FullPlane(lilLen)
	} else if length-1 >= diff {
		result = FullPlane(lilLen) + 5 + 3*(diff-1)
	} else {
		result = FullPlane(lilLen) + 10 + 3*(diff-2)
	}

	return result
}

func CubeEdge(exp, length int) int {
	switch {
	case exp <= length:
		return 5 + (exp-1)*3
	case exp <= 2*length:
		return 10 + (exp-2)*3
	case exp <= 3*length:
		return 15 + (exp-3)*3
	}
	return 0
}

func Merging(exp int) int {
	var (
		matches              int
		lowLength, lowAmount = UpLow3(exp, "lower")
		diff                 = exp - lowAmount
		lilArea              = lowLength * lowLength
		lilDiff              = int(math.Floor(float64(diff) / float64(lilArea)))
	)

	//fmt.Println("lowAmount ", lowAmount)
	//fmt.Println("diff ", diff)
	//fmt.Println("lilArea ", lilArea)
	//fmt.Println("bigDiff", int(math.Ceil(float64(diff)/float64(lilArea))))
	//fmt.Println("lilDiff ", lilDiff)
	fmt.Println("cubes on one side", diff-lilDiff*lilArea)

	switch {
	case diff == 0:
		matches = FullKube(exp)
	case diff-lilDiff*lilArea == 0 && lilDiff <= 3:
		//fmt.Println("Full palne")
		// add int FullKube(lowAmount) +
		matches = FullKube(lowAmount) + FullPlane(lowLength)*lilDiff
	case diff-lilDiff*lilArea > 0 && math.Ceil(float64(diff)/float64(lilArea)) <= 3:
		fmt.Println("Plane")
		//fmt.Println(FullPlane(lowLength) * lilDiff)
		//fmt.Println(LilPlane(diff-lilDiff*lilArea, lowLength))
		matches = FullKube(lowAmount) + FullPlane(lowLength)*lilDiff + LilPlane(diff-lilDiff*lilArea, lowLength)
	case diff > 3*lilArea:
		fmt.Println("Edge")
		matches = FullKube(lowAmount) + FullPlane(lowLength)*3 + CubeEdge(diff-3*lilArea, lowLength)
	}

	return matches
}

func main() {
	var (
		n = 20
	)

	//_, _ = fmt.Scanln(&n)
	for i := 17; i <= n; i++ {
		fmt.Println(Merging(i), i)
	}
}
