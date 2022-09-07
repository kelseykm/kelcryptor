package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

// VerifyFiles ensures that files passed are either regular files or directories
func VerifyFiles(files []string) {
	var nonExistentFiles, nonRegularNonDirFiles []string

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			nonExistentFiles = append(nonExistentFiles, file)
		} else if !fileInfo.Mode().IsRegular() && !fileInfo.Mode().IsDir() {
			nonRegularNonDirFiles = append(nonRegularNonDirFiles, file)
		}
	}

	if len(nonExistentFiles) != 0 {
		fmt.Printf(
			"%serror:%s %ssome files do not exist:%s\n",
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)

		for index, file := range nonExistentFiles {
			fmt.Printf(
				"  %v: %s%s%s\n",
				index+1,
				colour.BrownItalicised,
				file,
				colour.Normal,
			)

			if index+1 == len(nonExistentFiles) {
				fmt.Println()
			}
		}

		flag.Usage()
		os.Exit(2)

	} else if len(nonRegularNonDirFiles) != 0 {
		fmt.Printf(
			"%serror:%s %ssome files are not regular files:%s\n",
			colour.RedBold,
			colour.Normal,
			colour.WhiteBold,
			colour.Normal,
		)

		for index, file := range nonRegularNonDirFiles {
			fmt.Printf(
				"  %v: %s%s%s\n",
				index+1,
				colour.BrownItalicised,
				file,
				colour.Normal,
			)

			if index+1 == len(nonRegularNonDirFiles) {
				fmt.Println()
			}
		}

		flag.Usage()
		os.Exit(2)
	}
}
