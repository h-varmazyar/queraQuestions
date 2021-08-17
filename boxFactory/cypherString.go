//question url:https://quera.ir/problemset/contest/106796/%D8%B3%D8%A4%D8%A7%D9%84-%D8%B1%D8%B4%D8%AA%D9%87%D9%87%D8%A7-%D8%B1%D8%B4%D8%AA%D9%87%DB%8C-%D8%B1%D9%85%D8%B2%DB%8C
package boxFactory

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

//var strArr []byte

var input string

func cypherString() {
	reader := bufio.NewReader(os.Stdin)
	strLength, _ := reader.ReadString('\n')
	execCount, _ := reader.ReadString('\n')
	count, err := strconv.Atoi(strings.TrimSpace(execCount))
	if err != nil {
		return
	}
	length, err := strconv.Atoi(strings.TrimSpace(strLength))
	if err != nil {
		return
	}
	str := make([]byte, length)
	for i := 0; i < length; i++ {
		str[i], _ = reader.ReadByte()
	}
	input = string(str)
	for i := 0; i < count; i++ {
		cypheringString()
	}
	fmt.Println(input)
}

func cypheringString() {
	end := len(input) - 1

	tmp := []byte(input)

	first := tmp[end]
	for i := len(tmp) - 1; i > 0; i-- {
		tmp[i] = tmp[i-1]
	}
	tmp[0] = first

	for i, _ := range tmp {
		if tmp[i] == 'z' {
			tmp[i] = 'a'
		} else {
			tmp[i] = tmp[i] + 1
		}
	}
	input = string(tmp)
}
