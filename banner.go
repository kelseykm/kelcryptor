package main

import (
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

var banner = []byte{32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 46, 45, 34, 34, 45, 46, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 47, 32, 46, 45, 45, 46, 32, 92, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 47, 32, 47, 32, 32, 32, 32, 92, 32, 92, 10, 32, 95, 32, 32, 95, 95, 32, 32, 32, 32, 95, 32, 32, 95, 95, 95, 95, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 95, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 124, 32, 124, 32, 32, 32, 32, 124, 32, 124, 10, 124, 32, 124, 47, 32, 47, 95, 95, 95, 124, 32, 124, 47, 32, 95, 95, 95, 124, 95, 32, 95, 95, 32, 95, 32, 32, 32, 95, 32, 95, 32, 95, 95, 32, 124, 32, 124, 95, 32, 95, 95, 95, 32, 32, 95, 32, 95, 95, 32, 32, 32, 32, 124, 32, 124, 46, 45, 34, 34, 45, 46, 124, 10, 124, 32, 39, 32, 47, 47, 32, 95, 32, 92, 32, 124, 32, 124, 32, 32, 32, 124, 32, 39, 95, 95, 124, 32, 124, 32, 124, 32, 124, 32, 39, 95, 32, 92, 124, 32, 95, 95, 47, 32, 95, 32, 92, 124, 32, 39, 95, 95, 124, 32, 32, 47, 47, 47, 96, 46, 58, 58, 58, 58, 46, 96, 92, 10, 124, 32, 46, 32, 92, 32, 32, 95, 95, 47, 32, 124, 32, 124, 95, 95, 95, 124, 32, 124, 32, 32, 124, 32, 124, 95, 124, 32, 124, 32, 124, 95, 41, 32, 124, 32, 124, 124, 32, 40, 95, 41, 32, 124, 32, 124, 32, 32, 32, 32, 124, 124, 124, 32, 58, 58, 47, 32, 32, 92, 58, 58, 32, 59, 10, 124, 95, 124, 92, 95, 92, 95, 95, 95, 124, 95, 124, 92, 95, 95, 95, 95, 124, 95, 124, 32, 32, 32, 92, 95, 95, 44, 32, 124, 32, 46, 95, 95, 47, 32, 92, 95, 95, 92, 95, 95, 95, 47, 124, 95, 124, 32, 32, 32, 32, 124, 124, 124, 32, 58, 58, 92, 95, 95, 47, 58, 58, 32, 59, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 124, 95, 95, 95, 47, 124, 95, 124, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 92, 92, 92, 32, 39, 58, 58, 58, 58, 39, 32, 47, 47, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 96, 61, 39, 58, 45, 46, 46, 45, 39, 96, 10}

func printBanner() {
	fmt.Fprintf(os.Stdout, "%s%s%s%s\n", colour.Green, colour.Bold, banner, colour.Normal)
}
