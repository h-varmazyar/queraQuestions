//question url: https://quera.ir/contest/assignments/32393/problems/108522
package hamCode

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func StringString() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	arr := strings.Split(strings.TrimSpace(input), " ")

	n, _ := strconv.Atoi(strings.TrimSpace(arr[0]))
	k, _ := strconv.Atoi(strings.TrimSpace(arr[1]))

	iw := make([]string, n)
	kw := make([]string, k)

	a := ""
	b := ""
	for i := 0; i < n; i++ {
		in, _ := reader.ReadString('\n')
		a += in
	}
	for i := 0; i < k; i++ {
		in, _ := reader.ReadString('\n')
		b += in
	}
	iw = strings.Split(strings.TrimSpace(a), "\n")
	kw = strings.Split(strings.TrimSpace(b), "\n")
	for _, kord := range kw {
		count := 0
		for _, iord := range iw {
			if match1(kord, iord) || match2(kord, iord) || match3(kord, iord) {
				count++
			}
		}
		fmt.Println(count)
	}
}

func match1(k, n string) bool {
	if len(k)+1 == len(n) {
		found := false
		for i := 0; i < len(k); i++ {
			if n[i] != k[i] {
				n = n[:i] + n[i+1:]
				found = true
				break
			}
			if !found && k == n[:len(n)-1] {
				return true
			}
		}
	} else if len(n)+1 == len(k) {
		found := false
		for i := 0; i < len(n); i++ {
			if n[i] != k[i] {
				k = k[:i] + k[i+1:]
				found = true
				break
			}
			if !found && n == k[:len(k)-1] {
				return true
			}
		}
	}
	return n == k
}

func match2(k, n string) bool {
	if len(n) == len(k) {
		found := 0
		for i := 0; i < len(n); i++ {
			if n[i] != k[i] {
				found++
			}
		}
		return found < 2
	} else {
		return false
	}
}

func match3(k, n string) bool {
	if strings.ToLower(k) == strings.ToLower(n) {
		return true
	}
	return false
}
