package arraysAndStrings

// Implement an algorithm to determine if a string has all unique characters
// (not using other data structures instead of arrays and strings)

func verifyIsUnique(testString string) bool {
	var array = make([]bool, 128)

	for _, char := range testString {
		var hash = int(char)

		if array[hash] {
			return false
		}

		array[hash] = true
	}
	return true
}

func verifyIsUniqueMySolution(testString string) bool {
	var letterChecker = make(map[int32]int)

	for _, char := range testString {
		letterChecker[char] += 1

		if letterChecker[char] > 1 {
			return false
		}
	}
	return true
}

func ExecuteIsUnique() {
	var testString = "abcdefAa"

	println(verifyIsUnique(testString))
	println(verifyIsUniqueMySolution(testString))
}
