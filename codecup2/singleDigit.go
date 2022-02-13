package codecup2

//question url: https://quera.org/problemset/3539

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SingleDigit() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	resp := int32(0)
	for {
		resp = sum(input)
		if resp < 10 {
			fmt.Println(resp)
			return
		}
		input = fmt.Sprint(resp)
	}
}

func sum(input string) int32 {
	sum := int32(0)
	for _, s := range strings.TrimSpace(input) {
		sum += s - 48
	}
	return sum
}
