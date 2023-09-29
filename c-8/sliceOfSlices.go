package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	i, j := 0, 0
	for i < rows {
		j = 0
		row := make([]int, 0)
		for j < cols {
			fmt.Printf("row %v, col %v \n", i, j)
			row = append(row, i*j)
			j++
		}
		matrix = append(matrix, row)
		i++
	}
	return matrix
}

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	fmt.Printf("%v", createMatrix(rows, cols))
	fmt.Printf("Done \n")
}

func main() {
	test(5, 5)
}
