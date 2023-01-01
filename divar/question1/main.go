package main

import (
	"fmt"
	"log"
)

var sum = 0

func main() {
	digitArr, err := prepareInputs()
	if err != nil {
		log.Fatalf("invalid inputs: %v", err)
	}

	sumArr := sumArray(digitArr)
	printOutput(sumArr)
}

func prepareInputs() ([]int, error) {
	length := 0
	_, _ = fmt.Scan(&length)

	digitArr := make([]int, 0)
	for i := 0; i < length; i++ {
		num := 0
		_, _ = fmt.Scan(&num)
		_, _ = fmt.Scan()
		digitArr = append(digitArr, num)
		sum += num
	}
	return digitArr, nil
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
