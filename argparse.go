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
var ignoreFlag bool
var versionFlag bool

func parseFlags() (ignore, timeTaken, encrypt, decrypt bool, files []string) {
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"%s %s\n",
			colour.Info,
			colour.Message("v"+version),
		)
		os.Exit(0)
	}

	timeTaken = timeTakenFlag
	encrypt = encryptFlag
	decrypt = decryptFlag
	ignore = ignoreFlag
	files = flag.Args()

	var errorString string

	if len(files) == 0 {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error,
			colour.Message("file(s) to be decrypted/encrypted not provided\n\n"),
		)
		fmt.Print(errorString)
		flag.Usage()
		os.Exit(2)
	} else if encryptFlag && decryptFlag {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error,
			colour.Message("cannot set both -decrypt and -encrypt\n\n"),
		)
		fmt.Print(errorString)
		flag.Usage()
		os.Exit(2)
	} else if !encryptFlag && !decryptFlag {
		errorString = fmt.Sprintf(
			"%s %s",
			colour.Error,
			colour.Message("either -decrypt or -encrypt must be set\n\n"),
		)
		fmt.Print(errorString)
		flag.Usage()
		os.Exit(2)
	}
	return
}

func init() {
	var flags []string
	var nonFlags []string

	for _, arg := range os.Args[1:] {
		if arg[0] == '-' {
			flags = append(flags, arg)
		} else {
			nonFlags = append(nonFlags, arg)
		}
	}

	args := append([]string{os.Args[0]}, flags...)
	args = append(args, nonFlags...)

	os.Args = args

	flag.Usage = func() {
		fmt.Print(
			"usage: kelcryptor [-h|-help] [-v|-version] [-i|-ignore] [[-e|-encrypt]|[-d|-decrypt]] [-t|-time] FILE [FILE ...]\n\n",
		)
		flag.PrintDefaults()
	}

	flag.BoolVar(&versionFlag, "version", false, "show version and exit")
	flag.BoolVar(&versionFlag, "v", false, "show version and exit (short option)")

	flag.BoolVar(&timeTakenFlag, "time", false, "show time taken to encrypt/decrypt file(s)")
	flag.BoolVar(&timeTakenFlag, "t", false, "show time taken to encrypt/decrypt file(s) (short option)")

	flag.BoolVar(&encryptFlag, "encrypt", false, "encrypt file(s)")
	flag.BoolVar(&encryptFlag, "e", false, "encrypt file(s) (short option)")

	flag.BoolVar(&decryptFlag, "decrypt", false, "decrypt file(s)")
	flag.BoolVar(&decryptFlag, "d", false, "decrypt file(s) (short option)")

	flag.BoolVar(&ignoreFlag, "ignore", false, "skip file(s) with errors when encrypting/decrypting")
	flag.BoolVar(&ignoreFlag, "i", false, "skip file(s) with errors when encrypting/decrypting (short option)")
}
