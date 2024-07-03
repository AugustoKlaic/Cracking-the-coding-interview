package arraysAndStrings

import utils "cracking-the-coding-interview/commons"

func zeroMatrix(matrix *[][]int) {
	var matrixLengthX = len(*matrix)
	var matrixLengthY = len((*matrix)[0])
	var foundZeroCoordinateX, foundZeroCoordinateY []int

	for i := 0; i < matrixLengthX; i++ {
		for j := 0; j < matrixLengthY; j++ {
			if (*matrix)[i][j] == 0 {
				foundZeroCoordinateX = append(foundZeroCoordinateX, i)
				foundZeroCoordinateY = append(foundZeroCoordinateY, j)
			}
		}
	}

	for i := 0; i < matrixLengthX; i++ {
		for j := 0; j < matrixLengthY; j++ {
			if utils.Contains(foundZeroCoordinateX, i) || utils.Contains(foundZeroCoordinateY, j) {
				(*matrix)[i][j] = 0
			}
		}
	}
}

func ExecuteZeroMatrix() {
	matrix := [][]int{{1, 2, 3}, {4, 0, 6}, {7, 8, 9}}
	zeroMatrix(&matrix)
	utils.PrintMatrix(&matrix)
}
