//question url: https://quera.ir/contest/assignments/32393/problems
package hamCode

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func WordWord() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	v := []rune{'a', 'e', 'i', 'o', 'u'}
	count := 0
	for _, c := range input {
		for _, r := range v {
			if r == c {
				count++
			}
		}
	}
	fmt.Println(math.Pow(2, float64(count)))
}
