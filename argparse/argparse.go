package argparse

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

var timeTakenFlag bool
var encryptFlag bool
var decryptFlag bool

// ParseFlags parses flags and arguments and  returns them
func ParseFlags() (timeTaken, encrypt, decrypt bool, files []string) {
	flag.Parse()
	var errorString string

	if len(flag.Args()) == 0 {
		errorString = fmt.Sprintf(
			"%s%serror%s %sfile(s) to be decrypted/encrypted not provided%s\n\n",
			colour.RedBlinking,
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	} else if encryptFlag && decryptFlag {
		errorString = fmt.Sprintf(
			"%s%serror%s %scannot set both -decrypt and -encrypt%s\n\n",
			colour.RedBlinking,
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	} else if !encryptFlag && !decryptFlag {
		errorString = fmt.Sprintf(
			"%s%serror%s %seither -decrypt or -encrypt must be set%s\n\n",
			colour.RedBlinking,
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	}

	timeTaken = timeTakenFlag
	encrypt = encryptFlag
	decrypt = decryptFlag
	files = flag.Args()

	return
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: kelcryptor [-h|-help] [[-e|-encrypt]|[-d|-decrypt]] [-time] FILE [FILE ...]\n\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&timeTakenFlag, "time", false, "show time taken to encrypt/decrypt file(s)")
	flag.BoolVar(&timeTakenFlag, "t", false, "show time taken to encrypt/decrypt file(s) (short option)")
	flag.BoolVar(&encryptFlag, "encrypt", false, "encrypt file(s)")
	flag.BoolVar(&encryptFlag, "e", false, "encrypt file(s) (short option)")
	flag.BoolVar(&decryptFlag, "decrypt", false, "decrypt file(s)")
	flag.BoolVar(&decryptFlag, "d", false, "decrypt file(s) (short option)")
}
