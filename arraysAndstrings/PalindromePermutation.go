package arraysAndStrings

import (
	utils "github.com/AugustoKlaic/Cracking-the-coding-interview/commons"
	"strings"
)

// Given a string, write a function to check if it is a permutation of a palindrome

func checkPalindromePermutation(palindromePermutation []string) (bool, []string) {
	var permutationPalindromeList []string
	for _, permutedWord := range palindromePermutation {
		if checkPalindrome(permutedWord) {
			permutationPalindromeList = append(permutationPalindromeList, permutedWord)
		}
	}

	return len(permutationPalindromeList) != 0, permutationPalindromeList
}

func checkPalindrome(word string) bool {
	var reversedWord string
	for charIndex := len(word) - 1; charIndex >= 0; charIndex-- {
		reversedWord += string(word[charIndex])
	}
	return word == reversedWord
}

func ExecutePalindromePermutation() {
	var stringPalindromePermutation = "sasa"
	var permutedWords []string = utils.CreatePermutations(stringPalindromePermutation, 0, len(stringPalindromePermutation)-1)
	hasPalindromes, palindromeList := checkPalindromePermutation(permutedWords)

	println(hasPalindromes, strings.Join(palindromeList, ","))
}
