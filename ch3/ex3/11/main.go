package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "1234567890.0965"

	fmt.Println(comma(s))
}

func comma(str string) string {
	var buf bytes.Buffer
	parts := getParts(str)
	s := parts[0]
	floatPart := parts[1]
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

	if len(floatPart) > 0 {
		buf.WriteString(".")
		buf.WriteString(floatPart)
	}

	return buf.String()
}

func getParts(s string) []string {
	return strings.Split(s, ".")
}
