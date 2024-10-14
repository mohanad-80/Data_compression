package main

import "fmt"

func main() {
	testString := "ffcabracadabrarrarradff"
	result, alphabetDict := encode(testString)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString)
	fmt.Println("=====================================")
	testString2 := "rararbcrarbc"
	result, alphabetDict = encode(testString2)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString2)
	fmt.Println("=====================================")
	testString3 := "xyxyxyxyxyxyzzzzzzzz"
	result, alphabetDict = encode(testString3)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString3)
	fmt.Println("=====================================")
	testString4 := "wabba/wabba/wabba/wabba/woo/woo/woo"
	result, alphabetDict = encode(testString4)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString4)
	fmt.Println("=====================================")
	testString5 := "a/bar/array/by/barrayar/bay"
	result, alphabetDict = encode(testString5)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString5)
	fmt.Println("=====================================")
	testString6 := "shhi/has/hhi//h/hash/hahish/hahha"
	result, alphabetDict = encode(testString6)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	result, alphabetDict = []int{6, 3, 4, 5, 2, 3, 1, 6, 2, 9, 11, 16, 12, 14, 4, 20, 10, 8, 23, 13}, []string{"a", "/", "h", "i", "s", "t"}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString6)
	fmt.Println("=====================================")
	testString7 := "ratatatat/a/rat/at/a/rat"
	result, alphabetDict = encode(testString7)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	result, alphabetDict = []int{3, 1, 4, 6, 8, 4, 2, 1, 2, 5, 10, 6, 11, 13, 6}, []string{"a", "/", "r", "t"}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString7)
	fmt.Println("=====================================")
	testString8 := "barrayar/bar/by/barrayar/bay"
	result, alphabetDict = encode(testString8)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString8)
	fmt.Println("=====================================")
	testString9 := "THIS/IS/HIS/HIT"
	result, alphabetDict = encode(testString9)
	for _, char := range alphabetDict {
		fmt.Println(char)
	}
	for _, t := range result {
		fmt.Println(t)
	}
	fmt.Println("decoded :", decode(result, alphabetDict))
	result, alphabetDict = []int{4, 5, 3, 1, 2, 8, 2, 7, 9, 7, 4}, []string{"S", "/", "I", "T", "H"}
	fmt.Println("decoded :", decode(result, alphabetDict))
	fmt.Println("original:", testString9)
}

func encode(pattern string) ([]int, []string) {
	alphabetDict := createAlphabetDict(pattern)
	dict := alphabetDict
	output := []int{}
	currentPatternToCheck := string(pattern[0])
	latestFoundIdx := 0

	for i := 1; ; i++ {
		foundIdx := findIn(dict, currentPatternToCheck)

		if foundIdx == 0 {
			output = append(output, latestFoundIdx)
			dict = append(dict, currentPatternToCheck)
			currentPatternToCheck = string(pattern[i-1])
			i--
		} else if i == len(pattern) {
			// in the case were we reach the end of the pattern
			// and the currentPatternToCheck is in the dict so we
			// add its index to the output
			output = append(output, foundIdx)
			break
		} else {
			currentPatternToCheck += string(pattern[i])
			latestFoundIdx = foundIdx
		}
	}

	return output, alphabetDict
}

func createAlphabetDict(pattern string) []string {
	dict := []string{}

	for _, char := range pattern {
		if findIn(dict, string(char)) == 0 {
			dict = append(dict, string(char))
		}
	}

	return dict
}

func findIn(a []string, element string) int {
	for i, v := range a {
		if v == element {
			return i + 1
		}
	}
	return 0
}

func decode(output []int, alphabetDict []string) string {
	dict := alphabetDict
	result := alphabetDict[output[0]-1]
	previousStep := result
	currentStep := ""

	for i := 1; i < len(output); i++ {
		index := output[i] - 1
		if index >= len(dict) {
			// if the index is not known in the dict yet we construct the
			// unknown by concatenating the previous step and the first
			// symbol in the previous step
			currentStep = string(previousStep[0])
			result += previousStep + currentStep
			dict = append(dict, previousStep+string(currentStep[0]))
			previousStep += currentStep
		} else {
			currentStep = dict[index]
			result += currentStep
			dict = append(dict, previousStep+string(currentStep[0]))
			previousStep = currentStep
		}
	}

	return result
}
