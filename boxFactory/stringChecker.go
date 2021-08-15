//question url: https://quera.ir/problemset/contest/106795/%D8%B3%D8%A4%D8%A7%D9%84-%D9%BE%DB%8C%D8%A7%D8%AF%D9%87-%D8%B3%D8%A7%D8%B2%DB%8C-%D8%B1%D8%B4%D8%AA%D9%87%D9%87%D8%A7-%D8%B1%D8%B4%D8%AA%D9%87-%DA%86%DA%A9%D8%B1
package boxFactory

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func stringChecker() {
	reader := bufio.NewReader(os.Stdin)
	firstString, _ := reader.ReadString('\n')
	secondString, _ := reader.ReadString('\n')

	firstString=strings.TrimSpace(firstString)
	secondString=strings.TrimSpace(secondString)
	if firstString[0]==secondString[len(secondString)-1] {
		fmt.Println("YES")
	}else{
		fmt.Println("NO")
	}
}
