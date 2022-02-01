package main

import (
	"GoTheProgrammingLanguage/ch2/3/popcount"
	"crypto/sha256"
	"fmt"
	"math"
)

func main() {
	fata, morgana := "X", "x"
	c1 := sha256.Sum256([]byte(fata))
	c2 := sha256.Sum256([]byte(morgana))
	c1PopCount := popCount(c1[:])
	c2PopCount := popCount(c2[:])
	diff := math.Abs(float64(c1PopCount - c2PopCount))
	fmt.Println(diff)
}

func popCount(data []byte) int  {
	var popCount int

	for _, v := range data {
		bitsNumber := popcount.PopCountByRightMostBit(uint64(v))
		popCount += bitsNumber
	}

	return popCount
}


