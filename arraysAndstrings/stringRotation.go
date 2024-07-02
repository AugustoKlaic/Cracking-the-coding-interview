package arraysAndStrings

// Given a method "isSubstring" which checks if one word is a substring of another. Given two strings s1 and s2
// write a code to check if s2 is a rotation of s1 using only one call to isSubstring.

func isSubstring(string1, string2 string) bool {

	if len(string1) != len(string2) {
		return false
	}

	var string2Concat = string2 + string2
	var startOfOriginalWord int

	for i := 0; i < len(string2Concat); i++ {
		if string1[0] == string2Concat[i] {
			startOfOriginalWord = i
		}

		var indexEndOfOriginalWord = i + len(string1)
		if indexEndOfOriginalWord > len(string2Concat) {
			indexEndOfOriginalWord = len(string2Concat)
		}

		var foundPossibleWord = string2Concat[startOfOriginalWord:indexEndOfOriginalWord]

		if foundPossibleWord == string1 {
			return true
		}
	}

	return false
}

func ExecuteStringRotation() {
	var string1, string2 = "waterbkttle", "erbottlewat"

	println(isSubstring(string1, string2))
}
