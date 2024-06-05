package main

import "fmt"

// Given an image represented by N x N matrix, where each pixel in the image is represented by an integer, write a code
// to rotate the image by 90 degrees.

func rotate(matrix *[][]int) [][]int {

	var matrixLength = len(*matrix)
	var newMatrix [][]int = make([][]int, matrixLength)

	for i := 0; i < matrixLength; i++ {
		for j := 0; j < matrixLength; j++ {
			if newMatrix[j] == nil {
				newMatrix[j] = make([]int, matrixLength)
			}
			newMatrix[j][i] = (*matrix)[i][j]
		}
	}

	for k := 0; k < matrixLength; k++ {
		for l := 0; l < matrixLength/2; l++ {
			var swapperAux = newMatrix[k][l]
			var swapper2 = newMatrix[k][(matrixLength-1)-l]
			newMatrix[k][l] = swapper2
			newMatrix[k][(matrixLength-1)-l] = swapperAux
		}

	}

	return newMatrix
}

// linha 1 -> coluna N
// linha N -> coluna 1
// {1,2,3}  {1,4,7}  {7,4,1}
// {4,5,6}  {2,5,8}  {8,5,2}
// {7,8,9}  {3,6,9}  {9,6,3}
// linha vira coluna
// depois trocar colunas
func main() {

	matrix := [][]int{{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}

	var newMatrix = rotate(&matrix)

	for _, row := range newMatrix {
		for _, val := range row {
			fmt.Printf("%4d", val) // formata com 4 espaços para alinhar
		}
		fmt.Println() // nova linha após cada linha da matriz
	}
}
