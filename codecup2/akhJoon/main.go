//question url: https://quera.org/problemset/3538
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	dayMap := [7]bool{false, false, false, false, false, false, false}
	for i := 0; i < 3; i++ {
		_, _ = reader.ReadString('\n')
		input, _ := reader.ReadString('\n')
		days := strings.Split(input, " ")
		for _, day := range days {
			switch strings.TrimSpace(day) {
			case "shanbe":
				dayMap[0] = true
			case "1shanbe":
				dayMap[1] = true
			case "2shanbe":
				dayMap[2] = true
			case "3shanbe":
				dayMap[3] = true
			case "4shanbe":
				dayMap[4] = true
			case "5shanbe":
				dayMap[5] = true
			case "jome":
				dayMap[6] = true
			}
		}
	}
	count := 0
	for _, b := range dayMap {
		if !b {
			count++
		}
	}

	fmt.Println(count)
}
