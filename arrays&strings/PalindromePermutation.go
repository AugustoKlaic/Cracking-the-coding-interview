package main

// Given a string, write a function to check if it is a permutation of a palindrome

func checkPalindromePermutation(palindromePermutation string) {

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

func main() {
	var stringPalindromePermutation = "sasa"

	for _, word := range createPermutations(stringPalindromePermutation, 0, len(stringPalindromePermutation)-1) {
		println(word)
	}
}
