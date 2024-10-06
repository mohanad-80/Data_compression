package main

import (
	"fmt"
)

type tag struct {
	index    int
	codeword string // length will be 1
}

func main() {
	testString := "ffcabracadabrarrarradff"
	result := encode(testString)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println("decoded :", decode(result))
	fmt.Println("original:", testString)
	fmt.Println("=====================================")
	testString2 := "rararbcrarbc"
	result = encode(testString2)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println("decoded :", decode(result))
	fmt.Println("original:", testString2)
	fmt.Println("=====================================")
	testString3 := "xyxyxyxyxyxyzzzzzzzz"
	result = encode(testString3)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println("decoded :", decode(result))
	fmt.Println("original:", testString3)
	fmt.Println("=====================================")
	testString4 := "wabba/wabba/wabba/wabba/woo/woo/woo"
	result = encode(testString4)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println("decoded :", decode(result))
	fmt.Println("original:", testString4)
}

func encode(pattern string) []*tag {
	var tags []*tag
	var dictionary []string
	currentWord := ""
	latestFoundIdx := 0

	for i := 0; i < len(pattern); i++ {
		currentWord += string(pattern[i])

		index := findIn(dictionary, currentWord)

		if index == 0 || i == len(pattern)-1 {
			dictionary = append(dictionary, currentWord)
			tags = append(tags, &tag{latestFoundIdx, currentWord[len(currentWord)-1:]})
			currentWord = ""
		}

		latestFoundIdx = index
	}

	return tags
}

func findIn(a []string, element string) int {
	for i, v := range a {
		if v == element {
			return i + 1
		}
	}
	return 0
}

func decode(tags []*tag) string {
	dictionary := []string{""}
	result := ""
	currentWord := ""

	for _, tag := range tags {
		currentWord = dictionary[tag.index] + tag.codeword
		result += currentWord
		dictionary = append(dictionary, currentWord)
	}

	return result
}
