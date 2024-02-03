package main

// Implement an algorithm to determine if a string has all unique characters
// (not using other data structures instead of arrays and strings)

func verifyIsUnique(testString string, array []bool) bool {
	for _, char := range testString {
		var hash = int(char)

		if array[hash] {
			return false
		}

		array[hash] = true
	}
	return true
}

func verifyIsUniqueMySolution(testString string, letterChecker map[int32]int) bool {
	for _, char := range testString {
		letterChecker[char] += 1

		if letterChecker[char] > 1 {
			return false
		}
	}
	return true
}

func main() {
	var testString = "abcdefAa"
	var array = make([]bool, 128)
	var letterChecker = make(map[int32]int)

	println(verifyIsUnique(testString, array))
	println(verifyIsUniqueMySolution(testString, letterChecker))
}
