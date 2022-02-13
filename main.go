//question url: https://quera.org/problemset/3539
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num, _ := strconv.Atoi(strings.TrimSpace(input))

	resp := "W"
	for i := 0; i < num; i++ {
		resp += "o"
	}
	fmt.Printf("%sw!", resp)
}
