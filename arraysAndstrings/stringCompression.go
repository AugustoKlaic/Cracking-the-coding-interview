package arraysAndStrings

import "strconv"

// implement a method to perform a basic string compression using the counts of repeated characters

func compressString(stringTobeCompressed string) string {
	var auxArray = []rune(stringTobeCompressed)
	var letterCounter = 1
	var compressedString = ""

	for index, char := range auxArray {
		if index+1 < len(auxArray) && char == auxArray[index+1] {
			letterCounter++
		} else {
			compressedString += string(char) + strconv.Itoa(letterCounter)
			letterCounter = 1
		}
	}
	if len(compressedString) < len(stringTobeCompressed) {
		return compressedString
	} else {
		return stringTobeCompressed
	}
}

func ExecuteStringCompression() {

	var stringTobeCompressed = "aabccccaaa"
	println(compressString(stringTobeCompressed)) // a2b1c4a3
}
