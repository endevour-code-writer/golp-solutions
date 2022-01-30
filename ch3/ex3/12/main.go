package main

import (
	"fmt"
)

func main() {
	s := "secure"
	compared := "rescue"

	fmt.Println(areAnagram(s, compared))
}

func areAnagram(s, compared string) bool {
	if s == compared {
		return false
	}

	sLen, comparedLen := len(s), len(compared)

	if sLen != comparedLen {
		return false
	}

	symbolNumberInS := countUniqueByteNumber(s)
	symbolNumberInCompared := countUniqueByteNumber(compared)

	for k, v := range symbolNumberInS {
		elem, ok := symbolNumberInCompared[k]

		if !ok || elem != v {
			return false
		}
	}

	return true
}

func countUniqueByteNumber(s string) map[byte]int {
	symbolNumberInS := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		bt := s[i]
		if _, ok := symbolNumberInS[bt]; !ok {
			symbolNumberInS[bt] = 1
			continue
		}

		symbolNumberInS[bt]++

	}

	return symbolNumberInS
}
