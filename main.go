package main

import (
	"fmt"

	"github.com/kelseykm/kelcryptor/argparse"
	"github.com/kelseykm/kelcryptor/banner"
)

func main() {
	banner.PrintBanner()

	// get user intentions
	toRecordTimeTaken, toEncrypt, toDecrypt, files := argparse.ParseFlags()

	fmt.Println(toRecordTimeTaken, toEncrypt, toDecrypt, files)
}
