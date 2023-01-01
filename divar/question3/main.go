package main

import (
	"fmt"
)

var sum = 0

func main() {
	str1, str2 := prepareInputs()

	if len(str1) != len(str2) {
		fmt.Println("NO")
		return
	}

	m1 := make(map[uint8]int)
	m2 := make(map[uint8]int)

	for i := 0; i < len(str1); i++ {
		m1[str1[i]]++
	}
	for i := 0; i < len(str2); i++ {
		m2[str1[i]]++
	}

	if len(m1) != len(m2) {
		fmt.Println("NO")
		return
	}

	mr := make(map[uint8]struct{})

	for _, v2 := range m2 {
		for k1, v1 := range m1 {
			if v1 == v2 {
				mr[k1] = struct{}{}
				delete(m2, k1)
				break
			}
		}
	}

	if len(mr) == len(m1) {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")

}

func prepareInputs() (string, string) {
	str1 := ""
	str2 := ""
	_, _ = fmt.Scan(&str1)
	_, _ = fmt.Scan(&str2)

	return str1, str2
}

func sumArray(digitArr []int) []int {
	sumArr := make([]int, len(digitArr))
	for i := 0; i < len(digitArr); i++ {
		sumArr[i] = sum - digitArr[i]
	}
	return sumArr
}

func printOutput(sumArr []int) {
	out := fmt.Sprint(sumArr)
	fmt.Println(out[1 : len(out)-1])
}
