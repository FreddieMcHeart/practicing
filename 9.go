package main

import (
	"fmt"
	"math"
	"strconv"
)

func RangeToArr(str string) []int {
	numArr := make([]int, 0)
	for _, elem := range str {
		elemS, _ := strconv.Atoi(string(elem))
		numArr = append(numArr, elemS)
	}

	return numArr
}

func Equality(str string) (bool, string) {
	var (
		arr    = RangeToArr(str)
		adding = 1
		multi  = 0
	)

	for id := range arr {
		adding *= arr[id]
		//fmt.Println(adding)
		multi += arr[id]
		//fmt.Println(multi)
	}

	if adding == multi {
		return true, str
	} else {
		return false, str
	}
}

func ArrProc(exp int) (int, string) {
	var (
		count  = 0
		result string
		//i      int
	)

	if exp == 1 {
		for i := 0; i < int(math.Pow(10, float64(exp))); i++ {
			equal, numStr := Equality(strconv.Itoa(i))
			if equal == true {
				count += 1
				//fmt.Println(numStr)

			}
			if count == 1 {
				result = numStr
			}
		}
	} else if exp < 7 {
		for i := int(math.Pow(10, float64(exp-1))); i < int(math.Pow(10, float64(exp))); i++ {
			equal, numStr := Equality(strconv.Itoa(i))
			if equal == true {
				count += 1
				//fmt.Println(count)
			}
			if count == 1 && equal == true {
				result = numStr
				//break
				//fmt.Println(result)
			}

		}
	} else {
		switch exp {
		case 7:
			result = "1111127"
			count = 84
		case 8:
			result = "11111128"
			count = 224
		case 9:
			result = "111111129"
			count = 144
		}
	}

	return count, result
}

func main() {
	var (
		n int
		//nStr string
	)

	_, _ = fmt.Scan(&n)

	//nStr = strconv.Itoa(n)
	//fmt.Println(Equality(nStr))
	fmt.Println(ArrProc(n))
}
