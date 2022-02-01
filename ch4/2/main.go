package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
)

const (
	SHA256 = "sha256"
	SHA384 = "sha384"
	SHA512 = "sha512"
)

func main() {
	checksumType := flag.String("c", SHA256, "Create checksum by by algorithm type")
	help := flag.Bool("help", false, "Displays this help")
	flag.Parse()

	if  *help {
		flag.Usage()
		return
	}

	chars := flag.Arg(0)

	if "" == chars {
		fmt.Println("Please, put string to create checksum")
		return
	}

	checksum := createCheckSum(checksumType, chars)
	fmt.Printf("%x\n", checksum)
}

func createCheckSum(chType *string, chars string) []byte {
	var checksum []byte
	data := []byte(chars)
	switch *chType {
		case SHA384:
			ch := sha512.Sum384(data)
			checksum = ch[:]
		case SHA512:
			ch := sha512.Sum512(data)
			checksum = ch[:]
		default:
			ch := sha512.Sum512_256(data)
			checksum = ch[:]
	}

	return checksum
}