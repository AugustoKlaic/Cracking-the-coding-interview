package main

// Given two strings, write a method to decide if one is permutation of the other

func isPermutation(first, second string) bool {
	return true
}

func isPermutationMySolution(first, second string) bool {
	if len(first) != len(second) {
		return false
	}

	var mapCharCountFirst = make(map[int32]int)
	var mapCharCountSecond = make(map[int32]int)

	for _, char := range first {
		mapCharCountFirst[char] += 1
	}

	for _, char := range second {
		mapCharCountSecond[char] += 1
	}

	for _, char := range first {
		if mapCharCountFirst[char] != mapCharCountSecond[char] {
			return false
		}
	}

	return true
}

func main() {

	var x, y = "abcd", "dbca"
	println(isPermutationMySolution(x, y))
}
