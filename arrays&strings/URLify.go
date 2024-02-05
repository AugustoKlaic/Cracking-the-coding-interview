package main

import "strings"

// Write a method to replace all spaces in a string with '%20'. Assume that the string has sufficient space at the end
// to hold the additional characters, and that the 'true' length of the string is given to you

func urlify(originalUrl []string, trueLength int) []string {
	var matchingCharacter = " "
	var numberOfSpaces = countSpaces(originalUrl, 0, trueLength, matchingCharacter)
	var newIndex = trueLength - 1 + numberOfSpaces*2

	for oldIndex := trueLength - 1; oldIndex >= 0; oldIndex -= 1 {
		if originalUrl[oldIndex] == matchingCharacter {
			originalUrl[newIndex] = "0"
			originalUrl[newIndex-1] = "2"
			originalUrl[newIndex-2] = "%"
			newIndex -= 3
		} else {
			originalUrl[newIndex] = originalUrl[oldIndex]
			newIndex -= 1
		}
	}
	return originalUrl
}

func countSpaces(url []string, startPosition int, length int, character string) int {
	var count int = 0
	for i := startPosition; i < length; i++ {
		if url[i] == character {
			count++
		}
	}
	return count
}

func urlifyMySlution(originalUrl []string) []string {
	var spaceReplacer = "%20" // I assume this as one character... my mistake
	var matchingCharacter = " "

	var countSpaces = 0
	for index, character := range originalUrl {
		if character == matchingCharacter {
			countSpaces++
		} else {
			countSpaces = 0
		}

		if countSpaces > 1 {
			originalUrl[index] = ""
		} else if character == matchingCharacter {
			originalUrl[index] = spaceReplacer
		}
	}

	return originalUrl
}

func main() {
	var url = []string{"here", " ", "have", " ", "some", " ", " ", " ", "spaces", " ", " ", "in", " ", "the", " ", "url"}

	println(strings.Join(urlifyMySlution(url), ""))
	println(strings.Join(urlify(url, 13), ""))
}
