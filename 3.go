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
	//side length							   plane area with side length eq len
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

// CubeEdge returns matches on Cube edges
func CubeEdge(exp, length int) int {
	var numEdge = int(math.Ceil(float64(exp) / float64(length)))
	return 5*numEdge + (exp-numEdge)*3
}

// Merging main func that calculates total matches
func Merging(exp int) int {
	var (
		matches              int
		lowLength, lowAmount = UpLow3(exp, "lower")
		diff                 = exp - lowAmount
		lilArea              = lowLength * lowLength
		lilDiff              = int(math.Floor(float64(diff) / float64(lilArea)))
		freeCube             int
	)

	if exp == 4 {
		lilDiff = 2
	}

	freeCube = diff - lilDiff*lilArea

	switch {
	case diff == 0:
		matches = FullKube(exp)
	case lilDiff < 2 && freeCube == 0:
		matches = FullKube(lowAmount) + FullPlane(lowLength)*lilDiff
	case lilDiff == 2 && freeCube <= lowLength:
		matches = FullKube(lowAmount) + FullPlane(lowLength)*lilDiff + CubeEdge(freeCube, lowLength)
	case lilDiff <= 1:
		matches = FullKube(lowAmount) + FullPlane(lowLength)*lilDiff + LilPlane(freeCube, lowLength)
	case diff-lowLength-2*lilArea <= lilArea:
		matches = FullKube(lowAmount) + FullPlane(lowLength)*2 + CubeEdge(lowLength, lowLength) + LilPlane(diff-lowLength-2*lilArea, lowLength)
	case diff-lowLength-3*lilArea <= 2*lowLength:
		matches = FullKube(lowAmount) + FullPlane(lowLength)*3 + CubeEdge(diff-3*lilArea, lowLength)
	}

	return matches
}

func main() {
	var (
		n int
	)

	_, _ = fmt.Scanln(&n)
	fmt.Println(Merging(n))
}
