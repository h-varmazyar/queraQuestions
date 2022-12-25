package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	Value  int
	Row    int
	Column int
}

var reader = bufio.NewReader(os.Stdin)

func main() {
	input, _ := reader.ReadString('\n')

	count, _ := strconv.Atoi(strings.TrimSpace(input))

	matrix := createMatrix(count)

	minPoints, maxPoints := determineMinMax(matrix)

	extremum := determineExtremum(matrix, minPoints, maxPoints)

	printOutput(extremum)
}

func createMatrix(rowCount int) [][]int {
	matrix := make([][]int, rowCount)
	for i := 0; i < rowCount; i++ {
		row, _ := reader.ReadString('\n')
		items := strings.Split(row, " ")
		numItems := make([]int, len(items))
		for j, item := range items {
			num, _ := strconv.Atoi(strings.TrimSpace(item))
			numItems[j] = num
		}
		matrix[i] = numItems
	}
	return matrix
}

func determineMinMax(matrix [][]int) (minPoints []*Point, maxPoints []*Point) {
	intSize := 32 << (^uint(0) >> 63)
	for i := 0; i < len(matrix); i++ {
		min := &Point{Value: 1<<(intSize-1) - 1}
		max := &Point{Value: -1 << (intSize - 1)}
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] < min.Value {
				min.Value = matrix[i][j]
				min.Row = i
				min.Column = j
			}
			if matrix[i][j] > max.Value {
				max.Value = matrix[i][j]
				max.Row = i
				max.Column = j
			}
		}
		minPoints = append(minPoints, min)
		maxPoints = append(maxPoints, max)
	}
	return
}

func determineExtremum(matrix [][]int, minPoints, maxPoints []*Point) []*Point {
	extremum := make([]*Point, 0)
	for _, point := range minPoints {
		if isMinimumInColumn(matrix, point) {
			extremum = append(extremum, point)
			continue
		}
		if isMaximumInColumn(matrix, point) {
			extremum = append(extremum, point)
		}
	}

	for _, point := range maxPoints {
		if isMinimumInColumn(matrix, point) {
			extremum = append(extremum, point)
			continue
		}
		if isMaximumInColumn(matrix, point) {
			extremum = append(extremum, point)
		}
	}
	return extremum
}

func isMinimumInColumn(matrix [][]int, point *Point) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][point.Column] < point.Value {
			return false
		}
	}
	return true
}

func isMaximumInColumn(matrix [][]int, point *Point) bool {
	for i := 0; i < len(matrix); i++ {
		if matrix[i][point.Column] > point.Value {
			return false
		}
	}
	return true
}

func printOutput(extremum []*Point) {
	fmt.Println(len(extremum))
	for _, point := range extremum {
		fmt.Printf("%v %v\n", point.Column, point.Row)
	}
}
