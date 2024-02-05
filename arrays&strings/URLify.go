package main

import "strings"

// Write a method to replace all spaces in a string with '%20'. Assume that the string has sufficient space at the end
// to hold the additional characters, and that the 'true' length of the string is given to you

func urlifyMySlution(originalUrl []string) []string {
	var spaceReplacer = "%20"
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
}
