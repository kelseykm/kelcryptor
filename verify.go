package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

func verifyFiles(files []string) {
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
			"%s %s",
			colour.Error(),
			colour.Message("Some files do not exist:\n"),
		)

		for index, file := range nonExistentFiles {
			fmt.Printf(
				"  %v: %s\n",
				index+1,
				colour.FileName(file),
			)

			if index+1 == len(nonExistentFiles) {
				fmt.Println()
			}
		}

		flag.Usage()
		os.Exit(2)

	} else if len(nonRegularNonDirFiles) != 0 {
		fmt.Printf(
			"%s %s",
			colour.Error(),
			colour.Message("Some files are not regular files:\n"),
		)

		for index, file := range nonRegularNonDirFiles {
			fmt.Printf(
				"  %v: %s\n",
				index+1,
				colour.FileName(file),
			)

			if index+1 == len(nonRegularNonDirFiles) {
				fmt.Println()
			}
		}

		flag.Usage()
		os.Exit(2)
	}
}
