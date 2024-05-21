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
				if index > len(stringTwoArray)-1 {
					stringTwoArray = append(stringTwoArray, stringOneArray[index])
					countOperation++
				} else if stringOneArray[index] != stringTwoArray[index] {
					firstPart := append([]rune(nil), stringTwoArray[:index]...)
					secondPart := append([]rune(nil), stringTwoArray[index:]...)
					stringTwoArray = append(append(firstPart, stringOneArray[index]), secondPart...)
					countOperation++
				}
			}
		} else if len(stringOneArray) < len(stringTwoArray) {
			// removing from second string operation
			for index := 0; index < len(stringTwoArray); index++ {
				if index > len(stringOneArray)-1 {
					stringTwoArray = stringOneArray[:len(stringTwoArray)-1]
					countOperation++
				} else if stringOneArray[index] != stringTwoArray[index] {
					stringTwoArray = append(stringTwoArray[index:], stringTwoArray[index+1:]...)
					countOperation++
				}
			}
		} else {
			// replace operation
			for index := 0; index < len(stringOneArray); index++ {
				if stringOneArray[index] != stringTwoArray[index] {
					stringTwoArray[index] = stringOneArray[index]
					countOperation++
				}
			}
		}

		if countOperation <= 1 {
			return true
		} else {
			return false
		}
	}
}

func main() {

	println(checkEditsAway("pale", "ple"))   // = true
	println(checkEditsAway("pale", "pales")) // = true
	println(checkEditsAway("pale", "bale"))  // = true
	println(checkEditsAway("pale", "bake"))  // = false

}
