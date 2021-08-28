//question url: https://quera.ir/contest/assignments/32393/problems/108521
package hamCode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Soobasa() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	arr := strings.Split(strings.TrimSpace(input), " ")
	input, _ = reader.ReadString('\n')
	tmp := strings.Split(strings.TrimSpace(input), " ")

	n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
	a, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
	b, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
	if len(tmp) != n {
		return
	}

	goals := make([]int, n)

	for i, goal := range tmp {
		goals[i], _ = strconv.Atoi(strings.TrimSpace(goal))
	}

	first := make([]int, 0)
	second := make([]int, 0)
	for i := 0; i < len(goals)-1; i++ {
		if goals[i] > goals[i+1] {
			first = append(first, goals[:i+1]...)
			second = append(second, goals[i+1:]...)
			break
		}
	}
	if len(second) == 0 {
		second = append(second, goals...)
	}
	if second[len(second)-1] > 90+b {
		fmt.Println("NO")
		return
	}
	if len(first) > 0 && first[len(first)-1] > 45+a {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
}
