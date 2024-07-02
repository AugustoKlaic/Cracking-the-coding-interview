package commons

import "fmt"

// all the reusable functions will be added here

func PrintMatrix(matrix *[][]int) {
	for _, row := range *matrix {
		for _, val := range row {
			fmt.Printf("%4d", val) // formata com 4 espaços para alinhar
		}
		fmt.Println() // nova linha após cada linha da matriz
	}
}
