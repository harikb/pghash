package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/harikb/pghash/lib/pghash"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Bytes()
		h := pghash.HashAny(line)
		fmt.Printf("%d\n", h)
	}
}
