package arraysAndStrings

import "strings"

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

// add limit to permutation
func createPermutations(wordToPermute string, startOfString, sizeOfString int) []string {
	var permutedWords []string
	chars := []rune(wordToPermute)
	if startOfString == sizeOfString {
		permutedWords = append(permutedWords, string(chars))
	} else {
		for i := startOfString; i <= sizeOfString; i++ {
			chars[startOfString], chars[i] = chars[i], chars[startOfString]
			permutedWords = append(permutedWords, createPermutations(string(chars), startOfString+1, sizeOfString)...)
			chars[startOfString], chars[i] = chars[i], chars[startOfString]
		}
	}
	return permutedWords
}

func executePalindromePermutation() {
	var stringPalindromePermutation = "sasa"
	var permutedWords []string = createPermutations(stringPalindromePermutation, 0, len(stringPalindromePermutation)-1)
	hasPalindromes, palindromeList := checkPalindromePermutation(permutedWords)

	println(hasPalindromes, strings.Join(palindromeList, ","))

}
