package commons

import (
	"fmt"
	"sort"
	"strings"
)

// all the reusable functions will be added here

func PrintMatrix(matrix *[][]int) {
	for _, row := range *matrix {
		for _, val := range row {
			fmt.Printf("%4d", val) // formata com 4 espaços para alinhar
		}
		fmt.Println() // nova linha após cada linha da matriz
	}
}

func Contains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func MySort(receivedString string) string {
	var sorted []string = strings.Split(receivedString, "")
	sort.Slice(sorted, func(i int, j int) bool {
		return sorted[i] < sorted[j]
	})
	return strings.Join(sorted, "")
}

// CreatePermutations add limit to permutation
func CreatePermutations(wordToPermute string, startOfString, sizeOfString int) []string {
	var permutedWords []string
	chars := []rune(wordToPermute)
	if startOfString == sizeOfString {
		permutedWords = append(permutedWords, string(chars))
	} else {
		for i := startOfString; i <= sizeOfString; i++ {
			chars[startOfString], chars[i] = chars[i], chars[startOfString]
			permutedWords = append(permutedWords, CreatePermutations(string(chars), startOfString+1, sizeOfString)...)
			chars[startOfString], chars[i] = chars[i], chars[startOfString]
		}
	}
	return permutedWords
}
