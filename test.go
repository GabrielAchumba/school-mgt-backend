package main

import (
	"fmt"
	"math"
)

func IsPrimeNumber(x int) bool {
	if x < 2 {
		return false
	}
	sq_root := int(math.Sqrt(float64(x)))
	for i := 2; i <= sq_root; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func PrimeNumbers(x1, x2 int) []int {
	primeNumbers := make([]int, 0)
	if x1 < 2 || x2 < 2 {
		return primeNumbers
	}
	for x1 <= x2 {
		isPrime := IsPrimeNumber(x1)
		if isPrime {
			primeNumbers = append(primeNumbers, x1)
		}
		x1++
	}
	return primeNumbers
}

func main2() {
	lst := PrimeNumbers(0, 2)
	fmt.Print(lst)
}
