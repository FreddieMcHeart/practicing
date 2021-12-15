package main

import (
	"fmt"
	"math/big"
)

func Gap(g, m, n int) []int {
	tmp := make([]int, 0)

	for i := m; i <= n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			tmp = append(tmp, i)
		}
	}

	tmp = append(tmp, 0)
	result := make([]int, 0)
	result = nil

	for id := range tmp {
		if id+1 < len(tmp) && tmp[id+1]-tmp[id] == g {
			result = append(result, tmp[id])
			result = append(result, tmp[id]+g)
			break
		}
	}

	return result
}

func main() {
	fmt.Println(Gap(2, 100, 110))
	fmt.Println(Gap(4, 100, 110))
	fmt.Println(Gap(6, 100, 110))
	fmt.Println(Gap(8, 300, 400))
	fmt.Println(Gap(10, 300, 400))
	fmt.Println(Gap(6, 934040, 934200))

}
