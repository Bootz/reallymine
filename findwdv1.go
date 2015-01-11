// 11 january 2015
package main

import (
	"fmt"
	"os"
	"encoding/hex"
)

const fullsize = 1000204886016
const blocksize = 512

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n := int64(fullsize) - blocksize
	b := make([]byte, blocksize)

	for n > 0 {
		_, err := f.Seek(n, 0)
		if err != nil {
			panic(err)
		}
		nr, err := f.Read(b)
		if err != nil {
			panic(err)
		} else if nr != blocksize {
			panic(nr)
		}
		if b[0] == 'W' && b[1] == 'D' && b[2] == 'v' && b[3] == '1' {
			break
		}
		n -= blocksize
	}

	if n < 0 {
		fmt.Printf("no match\n")
		os.Exit(1)
	}

	fmt.Printf("%s\n", hex.Dump(b))
	fmt.Printf("%d\n", n)
}
