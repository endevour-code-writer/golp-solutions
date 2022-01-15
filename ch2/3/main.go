package main

import (
	"GoTheProgrammingLanguage/ch2/3/popcount"
	"fmt"
	"math"
)

func main()  {
	max := uint64(math.MaxUint64)
	pc := popcount.PopCount(max)
	fmt.Println(pc)
}