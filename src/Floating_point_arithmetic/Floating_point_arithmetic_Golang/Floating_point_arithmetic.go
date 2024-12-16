package main

import (
	"fmt"
)

type Range struct {
	low  float64
	high float64
}

func main() {
	// compress("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbcccccccccccccccccc")
	value, rangesTable := compress("acba")
	result := decompress(value, rangesTable, 4)
	fmt.Println(result)
}

func compress(s string) (float64, map[string]Range) {
	freqTable := buildFreqTable(s)
	rangesTable := buildRangesTable(freqTable, float64(len(s)))
	fmt.Println(freqTable)
	fmt.Println(rangesTable)

	var lower float64 = 0
	var upper float64 = 1
	var temp1, temp2 float64

	for i := 0; i < len(s); i++ {
		temp1 = lower + ((upper - lower) * rangesTable[string(s[i])].low)
		temp2 = lower + ((upper - lower) * rangesTable[string(s[i])].high)
		lower, upper = temp1, temp2
	}

	fmt.Println(lower)
	fmt.Println(upper)
	fmt.Println((upper + lower) / 2)
	return (upper + lower) / 2, rangesTable
}

func buildFreqTable(data string) map[string]float64 {
	table := make(map[string]float64)

	for _, char := range data {
		table[string(char)]++
	}

	return table
}

func buildRangesTable(freqData map[string]float64, dataSize float64) map[string]Range {
	rangesTable := make(map[string]Range)
	var cumulativeFreq float64

	for k, v := range freqData {
		rangesTable[k] = Range{cumulativeFreq / dataSize, ((cumulativeFreq + v) / dataSize)}
		cumulativeFreq += v
	}

	return rangesTable
}

func decompress(value float64, rangesTable map[string]Range, dataSize int) string {
	result := ""

	for len(result) < dataSize {
		symbol, rng := findValueRange(rangesTable, value)
		if symbol == "" {
			return ""
		}

		fmt.Println(value)
		fmt.Println(symbol)
		fmt.Println(rng)

		result += symbol
		value = (value - rng.low) / (rng.high - rng.low)
	}

	fmt.Println(result)
	return result
}

func findValueRange(rangesTable map[string]Range, value float64) (string, Range) {
	for k, v := range rangesTable {
		if value >= v.low && value <= v.high {
			return k, v
		}
	}
	return "", Range{-1, -1}
}
