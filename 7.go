package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Parse(RoStr string) (string, string) {
	var (
		first  string
		second string
	)

	first = strings.Split(RoStr, "+")[0]
	second = strings.Split(RoStr, "+")[1]

	return first, second
}

func RomaToDec(Roma string) int {
	var (
		result int
	)

	RomaToDecMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	// Crating array with numeric
	NumArr := make([]int, 0)
	for _, elem := range Roma {
		NumArr = append(NumArr, RomaToDecMap[string(elem)])
	}

	// Checking cases with IV IX XL... and sum result
	for id := range NumArr {
		if id+1 < len(NumArr) && NumArr[id] >= NumArr[id+1] {
			result = result + NumArr[id]
		}
		if id+1 < len(NumArr) && NumArr[id] < NumArr[id+1] {
			result = result - NumArr[id]
		}
	}
	/*
	   IV = 4  if i < i+1 { res = res - a[i] }
	   IX - 9
	   XL = 40
	   XLIV
	   MCDXLIX
	*/
	// Sum last element
	result = result + NumArr[len(NumArr)-1]

	return result
}

func Sum(exp string) int {

	first, second := Parse(exp)

	return RomaToDec(first) + RomaToDec(second)
}

func DecToRoma(exp int) string {
	var (
		result = ""
		expS   = strconv.Itoa(exp)
		lenS   = len(expS)
		revert int
	)

	DecToRomaMap := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
		4:    "IV",
		9:    "IX",
		40:   "XL",
		90:   "XC",
		400:  "CD",
		900:  "CM",
	}

	for i := 1; i <= lenS; i++ {
		revert = exp - exp%int(math.Pow(10, float64(lenS-i)))
		exp = exp % int(math.Pow(10, float64(lenS-i)))
		var (
			NumT    = 1
			counter = 0
		)

		// NumT == value contains count of 10, 100, 1000 digits, like 500 -> NumT == 5
		// count == powers of ten, like 100 -> 20^2 -> count == 2
		for i := revert; i >= 10; i /= 10 {
			NumT = i / 10
			counter += 1
		}

		// last cut numeric less than 10
		// revert == 4 -> NumT == 4
		if revert < 10 {
			NumT = revert
		}

		// Generating 1, 10, 100, 1000 numbers from current one
		countT := revert / NumT

		//check
		/*
		   1448
		   1000 / 1 == 1000; 1000*1
		   400 / 4 == 100;
		   40 / 4 == 10;
		*/

		result = result + DecToRomaMap[NumT*countT]
		// processing VI-VII like cases
		if NumT > 5 && NumT < 9 {
			result = result + DecToRomaMap[5*countT] + strings.Repeat(DecToRomaMap[1*countT], NumT-5)
			// processing I-III like cases
		} else if NumT > 1 && NumT < 4 {
			result = result + strings.Repeat(DecToRomaMap[1*countT], NumT)
		}
	}

	return result
}

func main() {
	var exp string

	_, _ = fmt.Scanln(&exp)

	fmt.Println(DecToRoma(Sum(exp)))
}
