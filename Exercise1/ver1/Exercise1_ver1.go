package main

import (
	"fmt"
	"math"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i = i + 1 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func Get_snt(Limit int) []int {
	var ans []int
	for i := 2; i <= Limit; i += 1 {
		if isPrime(i) {
			ans = append(ans, i)
		}
	}
	return ans
}
func main() {
	Limit := 20
	List_snt := Get_snt(Limit)
	for _, x := range List_snt {
		fmt.Println(x)
	}
}
