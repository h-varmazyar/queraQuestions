//question url: https://quera.ir/contest/assignments/32393/problems/108520
package hamCode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Paint() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	arr := strings.Split(strings.TrimSpace(input), " ")

	n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
	m, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
	k, _ := strconv.Atoi(strings.TrimSpace(arr[2]))

	in := make([]string, k)

	wall := make([][]int, n*m)
	for i := 0; i < k; i++ {
		input, _ = reader.ReadString('\n')
		in[i] = strings.TrimSpace(input)
	}

	for t, s := range in {
		tmp := strings.Split(s, " ")
		x, _ := strconv.Atoi(strings.TrimSpace(tmp[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(tmp[1]))
		z, _ := strconv.Atoi(strings.TrimSpace(tmp[2]))
		for i := x - 1; i < z+(x-1); i++ {
			for j := y - 1; j < z+(y-1); j++ {
				wall[m*i+j] = append(wall[m*i+j], t)
			}
		}
	}

	ma := make(map[string]int)

	for _, ints := range wall {
		key := fmt.Sprint(ints)
		if key != "[]" {
			ma[key] = 1
		}
	}
	fmt.Println(len(ma))
}
