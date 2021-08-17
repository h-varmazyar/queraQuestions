//question url:https://quera.ir/problemset/contest/106797/%D8%B3%D8%A4%D8%A7%D9%84-%D9%BE%DB%8C%D8%A7%D8%AF%D9%87-%D8%B3%D8%A7%D8%B2%DB%8C-%D8%A8%D8%B3%D8%AA%D9%87-%D8%AF%D8%B1-%D8%AC%D8%AF%D9%88%D9%84
package boxFactory

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func boxInMatrix() {
	reader := bufio.NewReader(os.Stdin)
	tmp, _ := reader.ReadString('\n')
	arr := strings.Split(tmp, " ")
	k, _ := strconv.Atoi(strings.TrimSpace(arr[2]))
	m, _ := strconv.Atoi(strings.TrimSpace(arr[1]))
	n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		lineArr := strings.Split(line, " ")
		tmpArr := make([]int, m)
		for j := 0; j < m; j++ {
			tmpArr[j], _ = strconv.Atoi(strings.TrimSpace(lineArr[j]))
		}
		matrix[i] = tmpArr
	}
	x, y := 0, 0
	for l := 0; l < k; l++ {
		if isPrime(matrix[x][y]) {
			if m%2 == 0 {
				y = ((m / 2) + 1 - ((y + 1) - m/2)) - 1
			} else {
				y = (m/2 - y) + m/2
			}
			if n%2 == 0 {
				x = ((n / 2) + 1 - ((x + 1) - n/2)) - 1
			} else {
				x = (n/2 - x) + n/2
			}
		} else {
			switch matrix[x][y] % 4 {
			case 0:
				if y++; y == m {
					y = 0
				}
			case 1:
				if x++; x == n {
					x = 0
				}
			case 2:
				if y--; y == -1 {
					y = m - 1
				}
			case 3:
				if x--; x == -1 {
					x = n - 1
				}
			}
		}
	}
	fmt.Printf("%d %d", x+1, y+1)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	sqRoot := int(math.Sqrt(float64(num)))
	for i := 2; i <= sqRoot; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
