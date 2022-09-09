package main

import (
	"fmt"
	"os"

	"github.com/kelseykm/kelcryptor/colour"
)

func verifyFiles(files []string) {
	var nonExistentFiles, nonRegularFiles []string

	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			nonExistentFiles = append(nonExistentFiles, file)
		} else if !fileInfo.Mode().IsRegular() {
			nonRegularFiles = append(nonRegularFiles, file)
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
		os.Exit(2)

	} else if len(nonRegularFiles) != 0 {
		fmt.Printf(
			"%s %s",
			colour.Error(),
			colour.Message("Some files are not regular files:\n"),
		)

		for index, file := range nonRegularFiles {
			fmt.Printf(
				"  %v: %s\n",
				index+1,
				colour.FileName(file),
			)

			if index+1 == len(nonRegularFiles) {
				fmt.Println()
			}
		}
		os.Exit(2)
	}
}
