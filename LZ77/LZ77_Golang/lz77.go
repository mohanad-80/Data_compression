package main

import "fmt"

type triple struct {
	offset, length int
	codeword       string // string length will be 1
}

var LookaheadSize = 6
var SearchWindowSize = 7

func findTheLongestMatch(p1, p2, lookaheadSize int, input *string) (int, int, string) {
	s := *input
	maxLength := 0
	offset := 0
	nextCharacter := string(s[p2])

	// looking for a match in the search window
	for i := p1; i < p2; i++ {
		// found match
		if s[i] == s[p2] {
			length := 1
			tempPtr := p2 + 1
			// check for more characters in the match
			for j := i + 1; ; j++ {
				// case were the pointer is in both the pattern and look-ahead buffers
				// AND the pointed at chars are equal
				if tempPtr < min(len(s)-1, p2+lookaheadSize) && s[j] == s[tempPtr] {
					length++
					tempPtr++
					continue
				}

				// case were the pointer is in the pattern buffer
				// AND the pointed at chars are equal
				if tempPtr < len(s)-1 && s[j] == s[tempPtr] {
					length++
				}

				// when they are not equal OR we hit the end of
				// the pattern or the look-ahead buffers
				if length >= maxLength {
					maxLength = length
					offset = p2 - i
					// we check if we are pass the last char in the buffer
					// we do not encode a char in the triple since the last
					// char is encoded in the offset and length.
					if tempPtr > len(s)-1 {
						nextCharacter = ""
					} else {
						nextCharacter = string(s[p2+length])
					}
				}
				break
			}
		}
	}

	return offset, maxLength, nextCharacter
}

func encode(s string) []*triple {
	var triples []*triple
	lookaheadPtr, searchPtr := 0, 0

	for lookaheadPtr < len(s) {
		offset, length, codeword := findTheLongestMatch(searchPtr, lookaheadPtr, LookaheadSize, &s) // use this if you want to have look-ahead limit
		// offset, length, codeword := findTheLongestMatch(searchPtr, lookaheadPtr, len(s)-lookaheadPtr, &s)

		triples = append(triples, &triple{offset: offset, length: length, codeword: codeword})

		lookaheadPtr += length + 1
		searchPtr += max(0, lookaheadPtr-SearchWindowSize) // use this if you want to have search limit
	}

	return triples
}

func decode(t []*triple) string {
	result := ""

	for _, v := range t {
		i, j := len(result)-v.offset, v.length
		for i >= 0 && j > 0 && len(result) > 0 {
			result += string(result[i])
			i, j = i+1, j-1
		}
		result += v.codeword
	}

	return result
}

func main() {
	testString := "ffcabracadabrarrarradff"
	result := encode(testString)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println(decode(result))
	fmt.Println("=====================================")
	testString2 := "rararbcrarbc"
	result = encode(testString2)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println(decode(result))
	fmt.Println("=====================================")
	testString3 := "xyxyxyxyxyxyzzzzzzzz"
	result = encode(testString3)
	for _, t := range result {
		fmt.Println(*t)
	}
	fmt.Println(decode(result))
}

// =====================
// match has more chars
// if s[j] == s[tempPtr] && tempPtr != min(len(s)-1) {
// 	length++
// 	tempPtr++
// } else if s[j] == s[tempPtr] && tempPtr == len(s)+1 {
// 	if length-1 >= maxLength {
// 		maxLength = length - 1
// 		offset = p2 - i
// 		nextCharacter = string(s[p2+length-1])
// 	}
// 	break
// } else {
// 	if length >= maxLength {
// 		maxLength = length
// 		offset = p2 - i
// 		nextCharacter = string(s[p2+length])
// 	}
// 	break
// }

// 	triples = append(triples, &triple{0, 0, string(s[0])})

// 	lookaheadPtr := 1
// 	searchWindowPtr := 0

// 	for i := 0; i < len(s); i++ {
// 		searchPtr := searchWindowPtr
// 		for j := searchPtr; j < searchWindow; j++ {
// 			if s[j] != s[lookaheadPtr] {
// 				continue
// 			}

// 			length, maxLength := 0, 0
// 			matchFound := false
// 			for k := j; !matchFound || k < lookaheadPtr; k++ {
// 				tempPtr := lookaheadPtr
// 				if s[k] != s[tempPtr] {
// 					tempPtr = lookaheadPtr
// 					maxLength = max(length, maxLength)
// 				} else {
// 					tempPtr++
// 					length++
// 					matchFound = true
// 				}
// 			}
// 		}
// 	}

// fmt.Println(findTheLongestMatch(0, 0, len(testString)-0, &testString))
// fmt.Println(findTheLongestMatch(0, 1, len(testString)-1, &testString))
// fmt.Println(findTheLongestMatch(0, 2, len(testString)-2, &testString))
// fmt.Println(findTheLongestMatch(0, 3, len(testString)-3, &testString))
// fmt.Println(findTheLongestMatch(0, 4, len(testString)-4, &testString))
// fmt.Println(findTheLongestMatch(0, 6, len(testString)-6, &testString))
// fmt.Println(findTheLongestMatch(0, 8, len(testString)-8, &testString))
// fmt.Println(findTheLongestMatch(0, 13, len(testString)-13, &testString))