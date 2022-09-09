package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

var timeTakenFlag bool
var encryptFlag bool
var decryptFlag bool

func parseFlags() (timeTaken, encrypt, decrypt bool, files []string) {
	flag.Parse()

	timeTaken = timeTakenFlag
	encrypt = encryptFlag
	decrypt = decryptFlag
	files = flag.Args()

	var errorString string

	if len(files) == 0 {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error(),
			colour.Message("file(s) to be decrypted/encrypted not provided\n\n"),
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	} else if encryptFlag && decryptFlag {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error(),
			colour.Message("cannot set both -decrypt and -encrypt\n\n"),
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	} else if !encryptFlag && !decryptFlag {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error(),
			colour.Message("either -decrypt or -encrypt must be set\n\n"),
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	}
	return
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"usage: kelcryptor [-h|-help] [[-e|-encrypt]|[-d|-decrypt]] [-t|-time] FILE [FILE ...]\n\n",
		)
		flag.PrintDefaults()
	}

	flag.BoolVar(&timeTakenFlag, "time", false, "show time taken to encrypt/decrypt file(s)")
	flag.BoolVar(&timeTakenFlag, "t", false, "show time taken to encrypt/decrypt file(s) (short option)")
	flag.BoolVar(&encryptFlag, "encrypt", false, "encrypt file(s)")
	flag.BoolVar(&encryptFlag, "e", false, "encrypt file(s) (short option)")
	flag.BoolVar(&decryptFlag, "decrypt", false, "decrypt file(s)")
	flag.BoolVar(&decryptFlag, "d", false, "decrypt file(s) (short option)")
}
