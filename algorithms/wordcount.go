package algorithms

import (
	"fmt"
	"slices"
	"strings"
)

func GetWordCount(input string) ([]string, map[string]int) {
	words := strings.Split(input, " ")
	var keys []string
	countMap := make(map[string]int)
	for _, word := range words {
		c := countMap[word]
		if c == 0 {
			keys = append(keys, word)
		}
		countMap[word] = c + 1
	}
	fmt.Println(countMap)
	fmt.Println(keys)
	slices.SortStableFunc(keys, func(i, j string) int {
		result := 0
		if countMap[i] > countMap[j] {
			result = -1
		}
		return result
	})
	return keys, countMap
}
