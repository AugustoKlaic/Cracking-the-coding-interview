package main

import "math"

func checkEditsAway(stringOne, stringTwo string) bool {
	if stringOne == stringTwo {
		return true
	} else if math.Abs(float64(len(stringOne))-float64(len(stringTwo))) > 1 {
		return false
	} else {
		var stringOneArray, stringTwoArray = []rune(stringOne), []rune(stringTwo)
		var countOperation = 0

		if len(stringOneArray) > len(stringTwoArray) {
			// adding to second string operation
			for index := 0; index < len(stringOneArray); index++ {
				if stringOneArray[index] != stringTwoArray[index] {
					firstPart := append([]rune(nil), stringTwoArray[:index]...)
					secondPart := append([]rune(nil), stringTwoArray[index:]...)
					stringTwoArray = append(append(firstPart, stringOneArray[index]), secondPart...)
					countOperation++
				}
			}
			println(string(stringOneArray), " - ", string(stringTwoArray))
		} else if len(stringOneArray) < len(stringTwoArray) {
			// removing from second string operation
			for index := 0; index < len(stringTwoArray); index++ {
				if stringOneArray[index] != stringTwoArray[index] {
					stringTwoArray = append(stringTwoArray[index:], stringTwoArray[index+1:]...)
					countOperation++
				}
			}
			println(string(stringOneArray), " - ", string(stringTwoArray))
		} else {
			// replace operation
			for index := 0; index < len(stringOneArray); index++ {
				if stringOneArray[index] != stringTwoArray[index] {
					stringTwoArray[index] = stringOneArray[index]
					countOperation++
				}
			}
			println(stringOneArray, " - ", stringTwoArray)
		}

		if countOperation <= 1 {
			return true
		} else {
			return false
		}
	}
}

// pale , ple = true
// pales, pale = true
// pale , bale = true
// pale , bake = false
func main() {

	var stringOne = "pale"
	var stringTwo = "ple"

	println(checkEditsAway(stringOne, stringTwo))
}
