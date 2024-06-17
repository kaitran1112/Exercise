package main

import "fmt"

func Get_snt_2(Limit int) []int {
	var ans []int
	var dd [100000]int
	for i := 2; i <= Limit; i += 1 {
		if dd[i] != 0 {
			continue
		}
		ans = append(ans, i)
		for j := i * i; j <= Limit; j += i {
			dd[j] = i
		}
	}
	return ans
}
func main() {
	Limit := 20
	List_snt := Get_snt_2(Limit)
	for _, x := range List_snt {
		fmt.Println(x)
	}
}
