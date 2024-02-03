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

func main() {
	var testString = "abcdefAa"
	var array = make([]bool, 128)

	println(verifyIsUnique(testString, array))
}
