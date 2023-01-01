package main

import (
	"fmt"
	"sync"
)

var (
	resp    = make(chan int, 100000000)
	end     = make(chan bool)
	wg      = new(sync.WaitGroup)
	intSize = 32 << (^uint(0) >> 63)
	max     = -1 << (intSize - 1)
)

func main() {
	length := 0
	_, _ = fmt.Scan(&length)

	go maxOfMax()

	digitArr := make([]int, 0)
	for i := 0; i < length; i++ {
		num := 0
		_, _ = fmt.Scan(&num)
		_, _ = fmt.Scan()

		go find(num, digitArr)
		digitArr = append(digitArr, num)
	}
	wg.Wait()
	end <- true
	fmt.Println(max)
}

func find(num int, digits []int) {
	wg.Add(1)
	defer wg.Done()
	m := -1 << (intSize - 1)
	for _, digit := range digits {
		diff := num - digit
		if diff > m {
			m = diff
		}
	}
	resp <- m
}

func maxOfMax() {
	for {
		select {
		case <-end:
			return
		case num := <-resp:
			if num > max {
				max = num
			}
		}
	}
}
