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

func verifyFiles(files []string) (nonExistentFiles, nonRegularFiles []string) {
	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			nonExistentFiles = append(nonExistentFiles, file)
		} else if !fileInfo.Mode().IsRegular() {
			nonRegularFiles = append(nonRegularFiles, file)
		}
	}

	return
}

// ParseFlags parses flags and arguments and  returns them
func ParseFlags() (timeTaken, encrypt, decrypt bool, files []string) {
	flag.Parse()

	timeTaken = timeTakenFlag
	encrypt = encryptFlag
	decrypt = decryptFlag
	files = flag.Args()

	var errorString string

	if len(files) == 0 {
		errorString = fmt.Sprintf(
			"%serror:%s %sfile(s) to be decrypted/encrypted not provided%s\n\n",
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
			"%serror:%s %scannot set both -decrypt and -encrypt%s\n\n",
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
			"%serror:%s %seither -decrypt or -encrypt must be set%s\n\n",
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)
		flag.Usage()
		os.Exit(2)
	}

	nonExistentFiles, nonRegularFiles := verifyFiles(files)
	if len(nonExistentFiles) != 0 {
		errorString = fmt.Sprintf(
			"%serror:%s %ssome files do not exist:%s\n",
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)

		for index, file := range nonExistentFiles {
			fmt.Printf(
				"  %v: %s%s%s\n",
				index+1,
				colour.BrownItalicised,
				file,
				colour.Normal,
			)
			if index+1 == len(nonRegularFiles) {
				fmt.Println()
			}
		}
		flag.Usage()
		os.Exit(2)
	} else if len(nonRegularFiles) != 0 {
		errorString = fmt.Sprintf(
			"%serror:%s %ssome files are not regular files:%s\n",
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)
		fmt.Fprintf(flag.CommandLine.Output(), errorString)

		for index, file := range nonRegularFiles {
			fmt.Printf(
				"  %v: %s%s%s\n",
				index+1,
				colour.BrownItalicised,
				file,
				colour.Normal,
			)
			if index+1 == len(nonRegularFiles) {
				fmt.Println()
			}
		}
		flag.Usage()
		os.Exit(2)
	}

	return
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "usage: kelcryptor [-h|-help] [[-e|-encrypt]|[-d|-decrypt]] [-t|-time] FILE [FILE ...]\n\n")
		flag.PrintDefaults()
	}

	flag.BoolVar(&timeTakenFlag, "time", false, "show time taken to encrypt/decrypt file(s)")
	flag.BoolVar(&timeTakenFlag, "t", false, "show time taken to encrypt/decrypt file(s) (short option)")
	flag.BoolVar(&encryptFlag, "encrypt", false, "encrypt file(s)")
	flag.BoolVar(&encryptFlag, "e", false, "encrypt file(s) (short option)")
	flag.BoolVar(&decryptFlag, "decrypt", false, "decrypt file(s)")
	flag.BoolVar(&decryptFlag, "d", false, "decrypt file(s) (short option)")
}
