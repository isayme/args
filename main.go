package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Printf("argc: %d, argv:\n", len(args))
	for idx, arg := range args {
		fmt.Printf("  %d: %s\n", idx+1, arg)
	}
}
