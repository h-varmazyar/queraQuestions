package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	Value  int
	Row    int
	Column int
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	input, _ := reader.ReadString('\n')

	number := calculateNumber(input)

	printOutput(number)
}

func calculateNumber(inputStr string) float64 {
	trimmed := trimSpace(inputStr)
	sign, noSigned := extractSign(trimmed)
	noSigned = appendZero(noSigned)
	if !isStartWithNumber(noSigned) {
		return 0
	}
	number := extractNumber(noSigned)
	return number * float64(sign)
}

func trimSpace(s string) string {
	trimmed := ""
	for start := 0; start < len(s); start++ {
		if isWithSpace(s[start]) {
			continue
		} else if !isDigitSign(s[start]) {
			trimmed += s[start:]
			return trimmed
		} else {
			trimmed = fmt.Sprintf("%v%c", trimmed, s[start])
		}
	}
	return s
}

func extractSign(str string) (int, string) {
	if str[0] == '-' {
		return -1, str[1:]
	}
	if str[0] == '+' {
		return 1, str[1:]
	}
	return 1, str
}

func appendZero(str string) string {
	if str[0] == '.' {
		return fmt.Sprintf("0.%v", str[1:])
	}
	return str
}

func isDigitChar(c uint8) bool {
	if c > 47 && c < 58 {
		return true
	}
	return false
}

func isDigitSign(c uint8) bool {
	if c == 46 || c == 45 || c == 43 { // 46 = . & 45= - & 43= +
		return true
	}
	return false
}

func isWithSpace(c uint8) bool {
	if rune(c) == ' ' {
		return true
	}
	return false
}

func isStartWithNumber(str string) bool {
	if isDigitChar(str[0]) {
		return true
	}
	return false
}

func extractNumber(str string) float64 {
	wholePart := float64(0)
	fractionalPart := float64(0)
	floatingPoint := false
	fractionalCoefficient := 0.1
	for i := 0; i < len(str); i++ {
		if !isDigitChar(str[i]) && str[i] != '.' {
			break
		}
		if floatingPoint {
			if isDigitChar(str[i]) {
				fractionalPart = fractionalPart*10 + float64(str[i]-48)
				fractionalCoefficient *= 0.1
			}
		} else {
			if isDigitChar(str[i]) {
				wholePart = wholePart*10 + float64(str[i]-48)
			}
		}
		if str[i] == '.' {
			floatingPoint = true
		}
	}
	num := float64(0)
	if floatingPoint {
		fractionalPart = fractionalPart * fractionalCoefficient / 0.1
		num = wholePart + fractionalPart
	} else {
		num = wholePart
	}
	return num
}

func printOutput(number float64) {
	fmt.Printf("%.3f", number*2)
}
