package main

import (
	"bufio"
	"os"

	"github.com/harikb/pghash/lib/pghash"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		h := pghash.HashAny(line)
	}
}
