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
		// not a match
		if s[i] != s[p2] {
			continue
		}
		// found match
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
			// but hit the end of the look-ahead buffer so we still
			// can count this char in the length
			// AND the pointed at chars are equal
			if tempPtr < len(s)-1 && s[j] == s[tempPtr] {
				length++
			}

			// if the match is not the longest do not count it
			if length < maxLength {
				break
			}

			// when the pointed at chars are not equal
			// OR we hit the end of the pattern or the look-ahead buffers
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
			break
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
