package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "1234567890"

	fmt.Println(comma(s))
}

func comma(s string) string {
	var buf bytes.Buffer
	l := len(s)

	if l <= 3 {
		return s
	}

	for i := 0; i < l; i++ {
		n := s[i:i+1]
		tail := s[i+1:l]
		tailLen := len(tail)
		remain := tailLen % 3
		buf.WriteString(n)

		if remain == 0 {
			buf.WriteString(",")
		}

		if tailLen == 3 {
			buf.WriteString(tail)
			break
		}
	}

	return buf.String()
}
