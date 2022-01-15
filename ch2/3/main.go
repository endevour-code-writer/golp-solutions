package main

import (
	"GoTheProgrammingLanguage/ch2/3/popcount"
	"fmt"
	"math"
)

func main()  {
	max := uint64(math.MaxUint64) - 100
	pc := popcount.PopCount(max)
	pcl := popcount.PopCountViaLoop(max)
	pcshift := popcount.PopCountByRightShift(max)
	pclmb := popcount.PopCountByRightShift(max)
	fmt.Println(pc)
	fmt.Println(pcl)
	fmt.Println(pcshift)
	fmt.Println(pclmb)
}